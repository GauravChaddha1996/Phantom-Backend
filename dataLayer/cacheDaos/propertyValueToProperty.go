package cacheDaos

import (
	"github.com/gomodule/redigo/redis"
	"phantom/dataLayer/dbModels"
	"strconv"
)

const PropertyValueIdToPropertyIdCacheName = "property_value_id_to_property_id_cache"

type PropertyValueIdToPropertyIdRedisDao struct {
	Pool *redis.Pool
}

func (dao PropertyValueIdToPropertyIdRedisDao) DeleteWholeCache(propertyValueArr *[]dbModels.PropertyValue) error {
	conn := dao.Pool.Get()
	for _, propertyValue := range *propertyValueArr {
		key := PropertyValueIdToPropertyIdCacheName + ":" + strconv.FormatInt(propertyValue.Id, 10)
		_, err := conn.Do("DEL", key)
		if err != nil {
			return err
		}
	}
	return nil
}

func (dao PropertyValueIdToPropertyIdRedisDao) SetPropertyValueIdToPropertyIdCache(dataArr *[]dbModels.PropertyValue) error {
	conn := dao.Pool.Get()
	for _, propertyValue := range *dataArr {
		key := PropertyValueIdToPropertyIdCacheName + ":" + strconv.FormatInt(propertyValue.Id, 10)
		_, err := conn.Do("SET", key, propertyValue.PropertyId)
		if err != nil {
			return err
		}
	}
	return nil
}

func (dao PropertyValueIdToPropertyIdRedisDao) ReadPropertyIdForPropertyValueId(propertyValueId int64) (*int64, error) {
	conn := dao.Pool.Get()
	key := PropertyValueIdToPropertyIdCacheName + ":" + strconv.FormatInt(propertyValueId, 10)
	propertyId, err := redis.Int64(conn.Do("GET", key))
	if err != nil {
		return nil, err
	}
	return &propertyId, nil
}
