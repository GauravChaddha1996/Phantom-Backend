package cacheDaos

import (
	"github.com/gomodule/redigo/redis"
	"github.com/spf13/cast"
)

const PropertyValueIdToProductIdCacheName = "property_value_id_to_product_id_cache"

type PropertyValueToProductRedisDao struct {
	Pool *redis.Pool
}

func (dao PropertyValueToProductRedisDao) GetCacheName(propertyId int64) string {
	return PropertyValueIdToProductIdCacheName + ":" + cast.ToString(propertyId)
}
