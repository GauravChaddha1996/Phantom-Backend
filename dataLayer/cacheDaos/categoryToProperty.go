package cacheDaos

import (
	"github.com/gomodule/redigo/redis"
	"github.com/hashicorp/go-multierror"
	"github.com/spf13/cast"
	"phantom/dataLayer/dbModels"
)

const CategoryIdToPropertyIdCacheName = "category_id_to_property_id_cache"

type CategoryIdToPropertyIdRedisDao struct {
	Pool *redis.Pool
}

func (dao CategoryIdToPropertyIdRedisDao) DeleteWholeCache(categoryArr *[]dbModels.Category) error {
	conn := dao.Pool.Get()
	for _, category := range *categoryArr {
		key := CategoryIdToPropertyIdCacheName + ":" + cast.ToString(category.Id)
		_, deleteErr := conn.Do("DEL", key)
		if deleteErr != nil {
			return deleteErr
		}
	}
	return nil
}

func (dao CategoryIdToPropertyIdRedisDao) SetCategoryToPropertyCache(dataArr *[]dbModels.CategoryToProperty) error {
	conn := dao.Pool.Get()
	var err error
	for _, categoryToProperty := range *dataArr {
		key := CategoryIdToPropertyIdCacheName + ":" + cast.ToString(categoryToProperty.CategoryId)
		_, categorySetErr := conn.Do("SADD", key, categoryToProperty.PropertyId)
		if categorySetErr != nil {
			err = multierror.Append(err, categorySetErr)
		}
	}
	return err
}

func (dao CategoryIdToPropertyIdRedisDao) ReadPropertyIdsForCategoryId(categoryId int64) (*[]int64, error) {
	conn := dao.Pool.Get()
	key := CategoryIdToPropertyIdCacheName + ":" + cast.ToString(categoryId)
	propertyIds, readErr := redis.Int64s(conn.Do("SMEMBERS", key))
	if readErr != nil {
		return nil, readErr
	}
	return &propertyIds, nil
}
