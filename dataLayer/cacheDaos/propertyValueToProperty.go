package cacheDaos

import (
	"github.com/gomodule/redigo/redis"
	"github.com/hashicorp/go-multierror"
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
		_, delErr := conn.Do("DEL", key)
		if delErr != nil {
			return delErr
		}
	}
	return nil
}

func (dao PropertyValueIdToPropertyIdRedisDao) SetPropertyValueIdToPropertyIdCache(dataArr *[]dbModels.PropertyValue) error {
	conn := dao.Pool.Get()
	var err error
	for _, propertyValue := range *dataArr {
		key := PropertyValueIdToPropertyIdCacheName + ":" + strconv.FormatInt(propertyValue.Id, 10)
		_, setErr := conn.Do("SET", key, propertyValue.PropertyId)
		if setErr != nil {
			err = multierror.Append(err, setErr)
		}
	}
	return err
}

func (dao PropertyValueIdToPropertyIdRedisDao) ReadPropertyIdForPropertyValueId(propertyValueId int64) (*int64, error) {
	conn := dao.Pool.Get()
	key := PropertyValueIdToPropertyIdCacheName + ":" + strconv.FormatInt(propertyValueId, 10)
	propertyId, readErr := redis.Int64(conn.Do("GET", key))
	if readErr != nil {
		return nil, readErr
	}
	return &propertyId, nil
}
