package home

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	"phantom/apis/apiCommons"
	"phantom/apis/home/models"
	"phantom/apis/home/section"
	"phantom/dataLayer/cacheDaos"
	"phantom/dataLayer/databasDaos"
	"phantom/dataLayer/dbModels"
	"phantom/dataLayer/uiModels/snippets"
	"phantom/ginRouter"
)

func ApiHandler(ctx *gin.Context) {
	// Initialize or find dependencies
	redisCachePool := ctx.MustGet(ginRouter.REDIS_POOL).(*redis.Pool)
	db := ctx.MustGet(ginRouter.SQL_DB).(*sql.DB)
	productCacheDao := &cacheDaos.AllProductIdsRedisDao{Pool: redisCachePool}
	categoryCacheDao := &cacheDaos.AllCategoryIdsRedisDao{Pool: redisCachePool}
	categoryToProductCacheDao := &cacheDaos.CategoryIdToProductIdRedisDao{Pool: redisCachePool}

	// Get db results
	apiDbResult := readFromDb(db)
	productIdMap := apiCommons.NewProductIdMap(apiDbResult.ProductsMap)

	// Make all sections
	newItemsProductRailSection := section.NewItemsProductRailSection(productCacheDao, productIdMap, apiDbResult)
	randomProductFullSection := section.RandomProductFullSection(productCacheDao, productIdMap, apiDbResult)
	categoryRailSection := section.CategoryRailSection(categoryCacheDao, apiDbResult)
	categoryTOProductRailSections := section.CategoryToProductRailSections(categoryToProductCacheDao, productIdMap, apiDbResult)
	randomProduct2FullSection := section.RandomProductFullSection(productCacheDao, productIdMap, apiDbResult)
	remainingDualProductSection := section.RemainingProductsSection(productIdMap.RemainingProducts(), apiDbResult)

	// Arrange sections
	var snippetSectionDataList []snippets.SnippetSectionData
	snippetSectionDataList = append(snippetSectionDataList, newItemsProductRailSection)
	snippetSectionDataList = append(snippetSectionDataList, randomProductFullSection)
	snippetSectionDataList = append(snippetSectionDataList, categoryRailSection)
	snippetSectionDataList = append(snippetSectionDataList, categoryTOProductRailSections...)
	snippetSectionDataList = append(snippetSectionDataList, randomProduct2FullSection)
	snippetSectionDataList = append(snippetSectionDataList, remainingDualProductSection)

	homeApiResponse := models.HomeApiResponse{
		Status:   "success",
		Message:  "",
		Snippets: snippetSectionDataList,
	}
	ctx.JSON(200, homeApiResponse)
}

func readFromDb(db *sql.DB) models.ApiDbResult {
	productDbDao := databasDaos.ProductSqlDao{DB: db}
	categoryDbDao := databasDaos.CategorySqlDao{DB: db}
	brandDbDao := databasDaos.BrandSqlDao{DB: db}

	productsChan := make(chan map[int64]*dbModels.Product, 1)
	categoriesChan := make(chan map[int64]*dbModels.Category, 1)
	brandsChan := make(chan map[int64]*dbModels.Brand, 1)

	go func(productsChan chan map[int64]*dbModels.Product) {
		products, _ := productDbDao.ReadAllProducts()
		productsChan <- mapProducts(products)
	}(productsChan)

	go func(categoriesChan chan map[int64]*dbModels.Category) {
		categories, _ := categoryDbDao.ReadAllCategories()
		categoriesChan <- mapCategories(categories)
	}(categoriesChan)

	go func(brandsChan chan map[int64]*dbModels.Brand) {
		brands, _ := brandDbDao.ReadAllBrands()
		brandsChan <- mapBrands(brands)
	}(brandsChan)

	return models.ApiDbResult{ProductsMap: <-productsChan, CategoriesMap: <-categoriesChan, BrandsMap: <-brandsChan}
}

func mapProducts(products *[]dbModels.Product) map[int64]*dbModels.Product {
	productMap := make(map[int64]*dbModels.Product, 0)
	for index := range *products {
		product := (*products)[index]
		productMap[product.Id] = &product
	}
	return productMap
}

func mapCategories(categories *[]dbModels.Category) map[int64]*dbModels.Category {
	categoryMap := make(map[int64]*dbModels.Category, 0)
	for index := range *categories {
		category := (*categories)[index]
		categoryMap[category.Id] = &category
	}
	return categoryMap
}

func mapBrands(brands *[]dbModels.Brand) map[int64]*dbModels.Brand {
	brandMap := make(map[int64]*dbModels.Brand, len(*brands))
	for index := range *brands {
		brand := (*brands)[index]
		brandMap[brand.Id] = &brand
	}
	return brandMap
}
