package cacheDaos

import (
	"github.com/gomodule/redigo/redis"
	"phantom/dataLayer/dbModels"
	"strconv"
)

const CategoryIdToPropertyIdCacheName = "category_id_to_property_id_cache"

type CategoryIdToPropertyIdRedisDao struct {
	Pool *redis.Pool
}

func (dao CategoryIdToPropertyIdRedisDao) DeleteWholeCache(categoryArr *[]dbModels.Category) error {
	conn := dao.Pool.Get()
	for _, category := range *categoryArr {
		key := CategoryIdToPropertyIdCacheName + ":" + strconv.FormatInt(category.Id, 10)
		_, deleteErr := conn.Do("DEL", key)
		if deleteErr != nil {
			return deleteErr
		}
	}
	return nil
}

func (dao CategoryIdToPropertyIdRedisDao) SetCategoryToPropertyCache(dataArr *[]dbModels.CategoryToProperty) error {
	conn := dao.Pool.Get()
	for _, categoryToProperty := range *dataArr {
		key := CategoryIdToPropertyIdCacheName + ":" + strconv.FormatInt(categoryToProperty.CategoryId, 10)
		_, err := conn.Do("SADD", key, categoryToProperty.PropertyId)
		if err != nil {
			return err
		}
	}
	return nil
}

func (dao CategoryIdToPropertyIdRedisDao) ReadPropertyIdsForCategoryId(categoryId int64) (*[]int64, error) {
	conn := dao.Pool.Get()
	key := CategoryIdToPropertyIdCacheName + ":" + strconv.FormatInt(categoryId, 10)
	propertyIds, err := redis.Int64s(conn.Do("SMEMBERS", key))
	if err != nil {
		return nil, err
	}
	return &propertyIds, nil
}
