package cacheDaos

import (
	"github.com/gomodule/redigo/redis"
	"phantom/dataLayer/dbModels"
	"strconv"
)

const PropertyIdToPropertyValueIdCacheName = "property_id_to_property_value_id_cache"

type PropertyIdToPropertyValueIdRedisDao struct {
	Pool *redis.Pool
}

func (dao PropertyIdToPropertyValueIdRedisDao) DeleteWholeCache(propertyArr *[]dbModels.Property) error {
	conn := dao.Pool.Get()
	for _, property := range *propertyArr {
		key := PropertyIdToPropertyValueIdCacheName + ":" + strconv.FormatInt(property.Id, 10)
		_, err := conn.Do("DEL", key)
		if err != nil {
			return err
		}
	}
	return nil
}

func (dao PropertyIdToPropertyValueIdRedisDao) SetPropertyIdToPropertyValueIdCache(dataArr *[]dbModels.PropertyValue) error {
	conn := dao.Pool.Get()
	for _, propertyValue := range *dataArr {
		key := PropertyIdToPropertyValueIdCacheName + ":" + strconv.FormatInt(propertyValue.PropertyId, 10)
		_, err := conn.Do("SADD", key, propertyValue.Id)
		if err != nil {
			return err
		}
	}
	return nil
}

func (dao PropertyIdToPropertyValueIdRedisDao) ReadPropertyValueIdsForPropertyId(propertyId int64) (*[]int64, error) {
	conn := dao.Pool.Get()
	key := PropertyIdToPropertyValueIdCacheName + ":" + strconv.FormatInt(propertyId, 10)
	propertyValueIdArr, err := redis.Int64s(conn.Do("SMEMBERS", key))
	if err != nil {
		return nil, err
	}
	return &propertyValueIdArr, nil
}
