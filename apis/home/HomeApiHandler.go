package home

import (
	"github.com/gomodule/redigo/redis"
	"log"
	"math/rand"
	"phantom/apis/apiCommons"
	"phantom/dataLayer/cacheDaos"
)

func ApiHandler(redisCachePool *redis.Pool) {
	log.Println("Starting home api handler")
	productCacheDao := &cacheDaos.AllProductIdsRedisDao{Pool: redisCachePool}
	categoryCacheDao := &cacheDaos.AllCategoryIdsRedisDao{Pool: redisCachePool}
	categoryToProductCacheDao := &cacheDaos.CategoryIdToProductIdRedisDao{Pool: redisCachePool}

	productIdMap := apiCommons.NewProductIdMap(productCacheDao)
	if productIdMap == nil {
		return
	}

	newItemsProductRailSectionIds := newItemsProductRailSection(productCacheDao, productIdMap)
	randomProductFullSectionId := randomProductFullSection(productCacheDao, productIdMap)
	categoryRailSectionIds := categoryRailSection(categoryCacheDao)
	categoryOneId, categoryTwoId := getTwoRandomCategoryId(categoryRailSectionIds)
	categoryOneRailSectionIds := categoryToProductRailSection(categoryOneId, categoryToProductCacheDao, productIdMap)
	categoryTwoRailSectionIds := categoryToProductRailSection(categoryTwoId, categoryToProductCacheDao, productIdMap)
	randomSecondProductFullSectionId := randomProductFullSection(productCacheDao, productIdMap)
	remainingDualProductSectionIds := productIdMap.RemainingProducts()

	log.Println("newItemsProductRailSectionIds: ", newItemsProductRailSectionIds)
	log.Println("randomProductFullSectionId: ", *randomProductFullSectionId)
	log.Println("categoryRailSectionIds: ", categoryRailSectionIds)
	log.Println("Category", *categoryOneId, ":", categoryOneRailSectionIds)
	log.Println("Category", *categoryTwoId, ":", categoryTwoRailSectionIds)
	if randomSecondProductFullSectionId != nil {
		log.Println("randomSecondProductFullSectionId: ", *randomSecondProductFullSectionId)
	}
	if remainingDualProductSectionIds != nil {
		log.Println("remainingDualProductSectionIds: ", remainingDualProductSectionIds)
	}
}

func newItemsProductRailSection(
	productCacheDao *cacheDaos.AllProductIdsRedisDao,
	productToFetchFromDbIds *apiCommons.ProductIdMap,
) *[]int64 {
	productRailSectionIds, err := productCacheDao.ReadFirstNProductIds(2)
	if err == nil {
		productToFetchFromDbIds.PutAllInt64s(productRailSectionIds)
	}
	return productRailSectionIds
}

func randomProductFullSection(
	productCacheDao *cacheDaos.AllProductIdsRedisDao,
	productToFetchFromDbIds *apiCommons.ProductIdMap,
) *int64 {
	shouldContinue := true
	totalIterations := 0
	for shouldContinue && totalIterations < 20 {
		totalIterations++
		randomProductId, err := productCacheDao.ReadRandomProduct()
		if err == nil {
			if productToFetchFromDbIds.Contains(*randomProductId) == false {
				shouldContinue = false
				productToFetchFromDbIds.Put(*randomProductId)
				return randomProductId
			}
		} else {
			shouldContinue = false
		}
	}
	return nil
}

func categoryRailSection(categoryDao *cacheDaos.AllCategoryIdsRedisDao) *[]int64 {
	allCategoryIds, err := categoryDao.ReadAllCategoryIds()
	if err != nil {
		return nil
	}
	return allCategoryIds
}

func getTwoRandomCategoryId(
	categoryIds *[]int64,
) (*int64, *int64) {
	maxSize := len(*categoryIds)
	indexOne := rand.Intn(maxSize)
	indexTwo := indexOne
	for indexTwo == indexOne {
		indexTwo = rand.Intn(maxSize)
	}
	categoryOne := (*categoryIds)[indexOne]
	categoryTwo := (*categoryIds)[indexTwo]
	return &categoryOne, &categoryTwo
}

func categoryToProductRailSection(
	categoryId *int64,
	categoryToProductCacheDao *cacheDaos.CategoryIdToProductIdRedisDao,
	productToFetchFromDbIds *apiCommons.ProductIdMap,
) *[]int64 {
	productsOfCategoryId, err := categoryToProductCacheDao.ReadNProductsOfCategoryId(categoryId, 1)
	if err != nil {
		return nil
	}
	productToFetchFromDbIds.PutAllInt64s(productsOfCategoryId)
	return productsOfCategoryId
}
