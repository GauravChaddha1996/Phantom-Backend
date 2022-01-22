package cacheDaos

import (
	"github.com/gomodule/redigo/redis"
	"github.com/hashicorp/go-uuid"
	"github.com/spf13/cast"
	"phantom/dataLayer/dbModels"
)

type FilterProductsDao struct {
	Pool *redis.Pool
}

func (dao FilterProductsDao) FindProductsForFilter(
	categoryId int64, propertyValueIdsMap map[int64]bool, allPropertyValues *[]dbModels.PropertyValue,
) (*[]int64, error) {
	conn := dao.Pool.Get()
	propertyValueToProductRedisDao := PropertyValueToProductRedisDao{dao.Pool}

	// Filtering is done like this
	// Intersection of
	// 		(category id)
	// 		(union of all applied properties of property id 1)
	// 		(union of all applied properties of property id 2)
	// 		....
	// This is done to serve the case
	// e.g. color selected are red and blue and fabric is cotton
	// The user intends to find either cotton red t-shirts or cotton blue t-shirts
	// So we have to take union of any property value id applied in its parent property id
	// But take intersection between those parent property ids

	// Find cache name for category id
	categoryIdToProductIdRedisDao := CategoryIdToProductIdRedisDao{dao.Pool}
	categoryIdCacheName := categoryIdToProductIdRedisDao.GetCacheName(categoryId)

	// Form all property value to property map
	allPropertyValueToPropertyMap := getAllPropertyValueToPropertyMap(allPropertyValues)

	// Form a map of applied property to array of applied property value cache name
	// e.g. appliedProperty is color = [blue,red]
	propertyToCacheNameMap := getPropertyToCacheNameMap(
		propertyValueIdsMap,
		propertyValueToProductRedisDao,
		allPropertyValueToPropertyMap,
	)

	// Take union of every applied property
	// e.g. for color union of blue and red
	// e.g. for fabric union of cotton and polyester
	unionResultsOfPropertyApplied, unionErr := unionResultsOfPropertyApplied(propertyToCacheNameMap, conn)
	if unionErr != nil {
		return nil, unionErr
	}

	// Set union results in a temp cache
	tempUnionResultCacheNames, tempUnionSetErr := setUnionResultInTempCache(unionResultsOfPropertyApplied, conn)
	if tempUnionSetErr != nil {
		return nil, tempUnionSetErr
	}

	// Take intersection of union results and find final filtered product ids
	finalCacheNameArr := combineAllCacheNames(categoryIdCacheName, *tempUnionResultCacheNames)
	productIds, intersectionErr := redis.Int64s(conn.Do("SINTER", finalCacheNameArr...))
	if intersectionErr != nil {
		return nil, intersectionErr
	}

	// Delete the temp cache we made of union results
	tempUnionCacheDelErr := deleteTempUnionResultCache(tempUnionResultCacheNames, conn)
	if tempUnionCacheDelErr != nil {
		return nil, tempUnionCacheDelErr
	}
	return &productIds, nil
}

func getAllPropertyValueToPropertyMap(allPropertyValues *[]dbModels.PropertyValue) map[int64]int64 {
	propertyValueToPropertyMap := map[int64]int64{}
	for _, propertyValueModel := range *allPropertyValues {
		propertyId := propertyValueModel.PropertyId
		propertyValueId := propertyValueModel.Id
		propertyValueToPropertyMap[propertyValueId] = propertyId
	}
	return propertyValueToPropertyMap
}

func getPropertyToCacheNameMap(
	propertyValueIdsMap map[int64]bool,
	propertyValueToProductRedisDao PropertyValueToProductRedisDao,
	allPropertyValueToPropertyMap map[int64]int64,
) map[int64][]string {
	propertyToCacheNameMap := map[int64][]string{}
	for propertyValueId := range propertyValueIdsMap {
		propertyValueCacheName := propertyValueToProductRedisDao.GetCacheName(propertyValueId)
		propertyId := allPropertyValueToPropertyMap[propertyValueId]
		var propertyIdCacheNameArr []string
		propertyIdCacheNameArr = append(propertyIdCacheNameArr, propertyToCacheNameMap[propertyId]...)
		propertyIdCacheNameArr = append(propertyIdCacheNameArr, propertyValueCacheName)
		propertyToCacheNameMap[propertyId] = propertyIdCacheNameArr
	}
	return propertyToCacheNameMap
}

func unionResultsOfPropertyApplied(
	propertyToCacheNameMap map[int64][]string,
	conn redis.Conn,
) (*map[int64][]int64, error) {
	propertyToProductIds := map[int64][]int64{}
	for propertyId := range propertyToCacheNameMap {
		unionCacheNames := combineAllCacheNames("", propertyToCacheNameMap[propertyId])
		unionProductIds, unionErr := redis.Int64s(conn.Do("SUNION", unionCacheNames...))
		if unionErr != nil {
			return nil, unionErr
		}
		propertyToProductIds[propertyId] = unionProductIds
	}
	return &propertyToProductIds, nil
}

func setUnionResultInTempCache(
	unionResultsOfPropertyApplied *map[int64][]int64,
	conn redis.Conn,
) (*[]string, error) {
	var intersectionCacheNames []string
	tempUUID, _ := uuid.GenerateUUID()
	for propertyId := range *unionResultsOfPropertyApplied {
		tempPropertyIdUnionResultCacheName := tempUUID + cast.ToString(propertyId)
		unionResultProductIds := (*unionResultsOfPropertyApplied)[propertyId]

		for _, productId := range unionResultProductIds {
			_, unionResultSetErr := conn.Do("SADD", tempPropertyIdUnionResultCacheName, productId)
			if unionResultSetErr != nil {
				return nil, unionResultSetErr
			}
		}
		intersectionCacheNames = append(intersectionCacheNames, tempPropertyIdUnionResultCacheName)
	}
	return &intersectionCacheNames, nil
}

func deleteTempUnionResultCache(tempUnionResultCacheNames *[]string, conn redis.Conn) error {
	for _, tempUnionResultCacheName := range *tempUnionResultCacheNames {
		_, unionResultDelErr := conn.Do("DEL", tempUnionResultCacheName)
		if unionResultDelErr != nil {
			return unionResultDelErr
		}
	}
	return nil
}

func combineAllCacheNames(
	categoryIdCacheName string, propertyValueCacheNameArr []string,
) []interface{} {
	var finalCacheNameArr []interface{}
	if len(categoryIdCacheName) != 0 {
		finalCacheNameArr = append(finalCacheNameArr, categoryIdCacheName)
	}
	for _, name := range propertyValueCacheNameArr {
		finalCacheNameArr = append(finalCacheNameArr, name)
	}
	return finalCacheNameArr
}
