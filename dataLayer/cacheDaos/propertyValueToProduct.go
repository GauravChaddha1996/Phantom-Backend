package cacheDaos

import (
	"github.com/gomodule/redigo/redis"
	"github.com/hashicorp/go-multierror"
	"phantom/dataLayer/dbModels"
	"strconv"
)

const PropertyValueIdToProductIdCacheName = "property_value_id_to_product_id_cache"

type PropertyValueToProductRedisDao struct {
	Pool *redis.Pool
}

func (dao PropertyValueToProductRedisDao) DeleteWholeCache(propertyValues *[]dbModels.PropertyValue) error {
	conn := dao.Pool.Get()

	for _, propertyValue := range *propertyValues {
		_, delErr := conn.Do("DEL", PropertyValueIdToProductIdCacheName+":"+strconv.FormatInt(propertyValue.Id, 10))
		if delErr != nil {
			return delErr
		}
	}
	return nil
}

func (dao PropertyValueToProductRedisDao) SetPropertyValuesToProductIds(dataArr *[]dbModels.ProductToProperty) error {
	conn := dao.Pool.Get()
	var err error
	for _, productToProperty := range *dataArr {
		key := PropertyValueIdToProductIdCacheName + ":" + strconv.FormatInt(productToProperty.ValueId, 10)
		_, setErr := conn.Do("SADD", key, productToProperty.ProductId)
		if setErr != nil {
			err = multierror.Append(err, setErr)
		}
	}
	return err
}

func (dao PropertyValueToProductRedisDao) ReadProductIdsOfPropertyValue(valueId int64) (*[]int64, error) {
	conn := dao.Pool.Get()
	key := PropertyValueIdToProductIdCacheName + ":" + strconv.FormatInt(valueId, 10)

	productIdsArr, cacheReadErr := redis.Int64s(conn.Do("SMEMBERS", key))
	if cacheReadErr != nil {
		return nil, cacheReadErr
	}
	return &productIdsArr, nil
}
