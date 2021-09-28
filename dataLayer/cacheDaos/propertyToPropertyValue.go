package cacheDaos

import (
	"github.com/gomodule/redigo/redis"
	"github.com/hashicorp/go-multierror"
	"github.com/spf13/cast"
	"phantom/dataLayer/dbModels"
)

const PropertyIdToPropertyValueIdCacheName = "property_id_to_property_value_id_cache"

type PropertyIdToPropertyValueIdRedisDao struct {
	Pool *redis.Pool
}

func (dao PropertyIdToPropertyValueIdRedisDao) DeleteWholeCache(propertyArr *[]dbModels.Property) error {
	conn := dao.Pool.Get()
	for _, property := range *propertyArr {
		key := PropertyIdToPropertyValueIdCacheName + ":" + cast.ToString(property.Id)
		_, delErr := conn.Do("DEL", key)
		if delErr != nil {
			return delErr
		}
	}
	return nil
}

func (dao PropertyIdToPropertyValueIdRedisDao) SetPropertyIdToPropertyValueIdCache(dataArr *[]dbModels.PropertyValue) error {
	conn := dao.Pool.Get()
	var err error
	for _, propertyValue := range *dataArr {
		key := PropertyIdToPropertyValueIdCacheName + ":" + cast.ToString(propertyValue.PropertyId)
		_, setErr := conn.Do("SADD", key, propertyValue.Id)
		if setErr != nil {
			err = multierror.Append(err, setErr)
		}
	}
	return err
}

func (dao PropertyIdToPropertyValueIdRedisDao) ReadPropertyValueIdsForPropertyId(propertyId int64) (*[]int64, error) {
	conn := dao.Pool.Get()
	key := PropertyIdToPropertyValueIdCacheName + ":" + cast.ToString(propertyId)
	propertyValueIdArr, readErr := redis.Int64s(conn.Do("SMEMBERS", key))
	if readErr != nil {
		return nil, readErr
	}
	return &propertyValueIdArr, nil
}
