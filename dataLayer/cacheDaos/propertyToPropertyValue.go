package cacheDaos

import (
	"github.com/gomodule/redigo/redis"
	"github.com/hashicorp/go-multierror"
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
		key := PropertyIdToPropertyValueIdCacheName + ":" + strconv.FormatInt(propertyValue.PropertyId, 10)
		_, setErr := conn.Do("SADD", key, propertyValue.Id)
		if setErr != nil {
			err = multierror.Append(err, setErr)
		}
	}
	return err
}

func (dao PropertyIdToPropertyValueIdRedisDao) ReadPropertyValueIdsForPropertyId(propertyId int64) (*[]int64, error) {
	conn := dao.Pool.Get()
	key := PropertyIdToPropertyValueIdCacheName + ":" + strconv.FormatInt(propertyId, 10)
	propertyValueIdArr, readErr := redis.Int64s(conn.Do("SMEMBERS", key))
	if readErr != nil {
		return nil, readErr
	}
	return &propertyValueIdArr, nil
}
