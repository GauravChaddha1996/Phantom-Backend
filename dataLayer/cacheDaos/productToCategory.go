package cacheDaos

import (
	"github.com/gomodule/redigo/redis"
	"phantom/dataLayer/dbModels"
	"strconv"
)

const CategoryIdToProductIdSetName = "category_id_to_product_id_set"

type CategoryIdToProductIdDao struct {
	Pool *redis.Pool
}

func (dao CategoryIdToProductIdDao) DeleteWholeCache(categoriesFromDb *[]dbModels.Category) error {
	conn := dao.Pool.Get()
	for _, category := range *categoriesFromDb {
		_, err := conn.Do("DEL", CategoryIdToProductIdSetName+":"+strconv.FormatInt(category.Id, 10))
		if err != nil {
			return err
		}
	}
	return nil
}

func (dao CategoryIdToProductIdDao) SetCategoryIdsToProductIdsMap(productArr *[]dbModels.Product) error {
	conn := dao.Pool.Get()
	for _, product := range *productArr {
		key := CategoryIdToProductIdSetName + ":" + strconv.FormatInt(product.CategoryId, 10)
		_, categoryIdSetErr := conn.Do("SADD", key, product.Id)
		if categoryIdSetErr != nil {
			return categoryIdSetErr
		}
	}
	return nil
}

func (dao CategoryIdToProductIdDao) ReadAllProductsOfCategoryId(categoryId int64) (*[]int64, error) {
	conn := dao.Pool.Get()
	key := CategoryIdToProductIdSetName + ":" + strconv.FormatInt(categoryId, 10)
	productIdStrArr, readCacheErr := redis.Strings(conn.Do("SMEMBERS", key))
	if readCacheErr != nil {
		return nil, readCacheErr
	}
	var productIdsArr = make([]int64, len(productIdStrArr))
	for index, productIdStr := range productIdStrArr {
		productId, atoiErr := strconv.ParseInt(productIdStr, 10, 64)
		if atoiErr != nil {
			return nil, atoiErr
		}
		productIdsArr[index] = productId
	}
	return &productIdsArr, nil
}
