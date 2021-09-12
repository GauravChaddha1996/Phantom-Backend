package cacheDaos

import (
	"github.com/gomodule/redigo/redis"
	"github.com/hashicorp/go-multierror"
	"phantom/dataLayer/dbModels"
)

const AllCategoryIdCacheName = "all_category_ids_cache"

type AllCategoryIdsRedisDao struct {
	Pool *redis.Pool
}

func (dao AllCategoryIdsRedisDao) DeleteWholeCache() error {
	conn := dao.Pool.Get()
	_, err := conn.Do("DEL", AllCategoryIdCacheName)
	if err != nil {
		return err
	}
	return nil
}

func (dao AllCategoryIdsRedisDao) SetCategoryIds(categoryArr *[]dbModels.Category) error {
	conn := dao.Pool.Get()
	var err error
	for _, category := range *categoryArr {
		_, categorySetErr := conn.Do("SADD", AllCategoryIdCacheName, category.Id)
		if categorySetErr != nil {
			err = multierror.Append(err, categorySetErr)
		}
	}
	return err
}

func (dao AllCategoryIdsRedisDao) ReadAllCategoryIds() (*[]int64, error) {
	conn := dao.Pool.Get()
	categoryIds, categoryIdsReadErr := redis.Int64s(conn.Do("SMEMBERS", AllCategoryIdCacheName))
	if categoryIdsReadErr != nil {
		return nil, categoryIdsReadErr
	}
	return &categoryIds, nil
}
