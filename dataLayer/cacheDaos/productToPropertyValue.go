package cacheDaos

import (
	"github.com/gomodule/redigo/redis"
	"phantom/dataLayer/dbModels"
	"strconv"
)

const ProductIdToPropertyValueIdSetName = "product_id_to_property_value_id"

type ProductToPropertyValueDao struct {
	Pool *redis.Pool
}

func (dao ProductToPropertyValueDao) DeleteWholeCache(products *[]dbModels.Product) error {
	conn := dao.Pool.Get()

	for _, product := range *products {
		_, err := conn.Do("DEL", ProductIdToPropertyValueIdSetName+":"+strconv.FormatInt(product.Id, 10))
		if err != nil {
			return err
		}
	}
	return nil
}

func (dao ProductToPropertyValueDao) SetProductIdsToPropertyValues(dataArr *[]dbModels.ProductToProperty) error {
	conn := dao.Pool.Get()
	for _, productToProperty := range *dataArr {
		key := ProductIdToPropertyValueIdSetName + ":" + strconv.FormatInt(productToProperty.ProductId, 10)
		_, err := conn.Do("SADD", key, productToProperty.ValueId)
		if err != nil {
			return err
		}
	}
	return nil
}

func (dao ProductToPropertyValueDao) ReadPropertyValueIdsOfProduct(productId int64) (*[]int64, error) {
	conn := dao.Pool.Get()
	key := ProductIdToPropertyValueIdSetName + ":" + strconv.FormatInt(productId, 10)

	propertyValueIdsArr, cacheReadErr := redis.Int64s(conn.Do("SMEMBERS", key))
	if cacheReadErr != nil {
		return nil, cacheReadErr
	}
	return &propertyValueIdsArr, nil
}
