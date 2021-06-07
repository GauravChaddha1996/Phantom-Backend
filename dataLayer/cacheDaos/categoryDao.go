package cacheDaos

import (
	"github.com/gomodule/redigo/redis"
	"phantom/dataLayer/dbModels"
)

const AllCategoryIdSetName = "all_category_ids_cache"

type AllCategoryIdsCacheDao struct {
	Pool *redis.Pool
}

func (dao AllCategoryIdsCacheDao) DeleteWholeCache() error {
	conn := dao.Pool.Get()
	_, err := conn.Do("DEL", AllCategoryIdSetName)
	if err != nil {
		return err
	}
	return nil
}

func (dao AllCategoryIdsCacheDao) SetCategoryIds(categoryArr *[]dbModels.Category) error {
	conn := dao.Pool.Get()
	for _, category := range *categoryArr {
		_, categorySetErr := conn.Do("SADD", AllCategoryIdSetName, category.Id)
		if categorySetErr != nil {
			return categorySetErr
		}
	}
	return nil
}

func (dao AllCategoryIdsCacheDao) ReadAllCategoryIds() (*[]string, error) {
	conn := dao.Pool.Get()
	categoryIds, categoryIdsReadErr := redis.Strings(conn.Do("SMEMBERS", AllCategoryIdSetName))
	if categoryIdsReadErr != nil {
		return nil, categoryIdsReadErr
	}
	return &categoryIds, nil
}
