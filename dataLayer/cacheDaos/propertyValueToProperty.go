package cacheDaos

import (
	"github.com/gomodule/redigo/redis"
	"github.com/hashicorp/go-multierror"
	"phantom/dataLayer/dbModels"
)

const AllPropertyValueIdCacheName = "all_property_value_id_cache"

type PropertyValueIdToPropertyIdRedisDao struct {
	Pool *redis.Pool
}

func (dao PropertyValueIdToPropertyIdRedisDao) DeleteWholeCache() error {
	conn := dao.Pool.Get()
	_, delErr := conn.Do("DEL", AllPropertyValueIdCacheName)
	if delErr != nil {
		return delErr
	}
	return nil
}

func (dao PropertyValueIdToPropertyIdRedisDao) SetPropertyValueIdsCache(dataArr *[]dbModels.PropertyValue) error {
	conn := dao.Pool.Get()
	var err error
	for _, propertyValue := range *dataArr {
		_, setErr := conn.Do("SADD", AllPropertyValueIdCacheName, propertyValue.Id)
		if setErr != nil {
			err = multierror.Append(err, setErr)
		}
	}
	return err
}

func (dao PropertyValueIdToPropertyIdRedisDao) IsPropertyIdValid(propertyValueId int64) (bool, error) {
	conn := dao.Pool.Get()
	exists, readErr := redis.Int(conn.Do("SISMEMBER", AllPropertyValueIdCacheName, propertyValueId))
	if readErr != nil {
		return false, readErr
	}
	return exists == 1, nil
}
