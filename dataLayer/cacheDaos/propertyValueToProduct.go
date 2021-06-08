package cacheDaos

import (
	"github.com/gomodule/redigo/redis"
	"phantom/dataLayer/dbModels"
	"strconv"
)

const PropertyValueIdToProductIdSetName = "property_value_id_to_product_id"

type PropertyValueToProductDao struct {
	Pool *redis.Pool
}

func (dao PropertyValueToProductDao) DeleteWholeCache(propertyValues *[]dbModels.PropertyValue) error {
	conn := dao.Pool.Get()

	for _, propertyValue := range *propertyValues {
		_, err := conn.Do("DEL", PropertyValueIdToProductIdSetName+":"+strconv.FormatInt(propertyValue.Id, 10))
		if err != nil {
			return err
		}
	}
	return nil
}

func (dao PropertyValueToProductDao) SetPropertyValuesToProductIds(dataArr *[]dbModels.ProductToProperty) error {
	conn := dao.Pool.Get()
	for _, productToProperty := range *dataArr {
		key := PropertyValueIdToProductIdSetName + ":" + strconv.FormatInt(productToProperty.ValueId, 10)
		_, err := conn.Do("SADD", key, productToProperty.ProductId)
		if err != nil {
			return err
		}
	}
	return nil
}

func (dao PropertyValueToProductDao) ReadProductIdsOfPropertyValue(valueId int64) (*[]int64, error) {
	conn := dao.Pool.Get()
	key := PropertyValueIdToProductIdSetName + ":" + strconv.FormatInt(valueId, 10)

	productIdsArr, cacheReadErr := redis.Int64s(conn.Do("SMEMBERS", key))
	if cacheReadErr != nil {
		return nil, cacheReadErr
	}
	return &productIdsArr, nil
}
