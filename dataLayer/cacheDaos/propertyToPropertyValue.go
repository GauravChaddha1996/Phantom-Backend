package cacheDaos

import (
	"github.com/gomodule/redigo/redis"
	"phantom/dataLayer/dbModels"
	"strconv"
)

const PropertyIdToPropertyValueIdSetName = "property_id_to_property_value_id"

type PropertyIdToPropertyValueIdDao struct {
	Pool *redis.Pool
}

func (dao PropertyIdToPropertyValueIdDao) DeleteWholeCache(propertyArr *[]dbModels.Property) error {
	conn := dao.Pool.Get()
	for _, property := range *propertyArr {
		key := PropertyIdToPropertyValueIdSetName + ":" + strconv.FormatInt(property.Id, 10)
		_, err := conn.Do("DEL", key)
		if err != nil {
			return err
		}
	}
	return nil
}

func (dao PropertyIdToPropertyValueIdDao) SetPropertyIdToPropertyValueIdCache(dataArr *[]dbModels.PropertyValue) error {
	conn := dao.Pool.Get()
	for _, propertyValue := range *dataArr {
		key := PropertyIdToPropertyValueIdSetName + ":" + strconv.FormatInt(propertyValue.PropertyId, 10)
		_, err := conn.Do("SADD", key, propertyValue.Id)
		if err != nil {
			return err
		}
	}
	return nil
}

func (dao PropertyIdToPropertyValueIdDao) ReadPropertyValueIdsForPropertyId(propertyId int64) (*[]int64, error) {
	conn := dao.Pool.Get()
	key := PropertyIdToPropertyValueIdSetName + ":" + strconv.FormatInt(propertyId, 10)
	propertyValueIdArr, err := redis.Int64s(conn.Do("SMEMBERS", key))
	if err != nil {
		return nil, err
	}
	return &propertyValueIdArr, nil
}
