package cacheDaos

import (
	"github.com/gomodule/redigo/redis"
	"github.com/hashicorp/go-multierror"
	"github.com/spf13/cast"
	"phantom/dataLayer/dbModels"
)

const PropertyValueIdToProductIdCacheName = "property_value_id_to_product_id_cache"

type PropertyValueToProductRedisDao struct {
	Pool *redis.Pool
}

func (dao PropertyValueToProductRedisDao) DeleteWholeCache(propertyValueIds *[]dbModels.PropertyValue) error {
	conn := dao.Pool.Get()
	for _, propertyValue := range *propertyValueIds {
		key := dao.GetCacheName(propertyValue.Id)
		_, deleteErr := conn.Do("DEL", key)
		if deleteErr != nil {
			return deleteErr
		}
	}
	return nil
}

func (dao PropertyValueToProductRedisDao) SetPropertyValueToProductCache(dataArr *[]dbModels.ProductToProperty) error {
	conn := dao.Pool.Get()
	var err error
	for _, productToProperty := range *dataArr {
		key := dao.GetCacheName(productToProperty.ValueId)
		_, propertyValueSetErr := conn.Do("SADD", key, productToProperty.ProductId)
		if propertyValueSetErr != nil {
			err = multierror.Append(err, propertyValueSetErr)
		}
	}
	return err
}

func (dao PropertyValueToProductRedisDao) GetCacheName(propertyId int64) string {
	return PropertyValueIdToProductIdCacheName + ":" + cast.ToString(propertyId)
}
