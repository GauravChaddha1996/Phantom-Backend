package home

import (
	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	"net/http"
	"phantom/apis/apiCommons"
	"phantom/apis/home/models"
	"phantom/apis/home/section"
	"phantom/dataLayer/cacheDaos"
	"phantom/dataLayer/uiModels/snippets"
	"phantom/ginRouter"
)

const Home_Db_Read_Err_Code = "Err code: 1"

func ApiHandler(ctx *gin.Context) {
	// Initialize or find dependencies
	redisCachePool := ctx.MustGet(ginRouter.REDIS_POOL).(*redis.Pool)
	productCacheDao := &cacheDaos.AllProductIdsRedisDao{Pool: redisCachePool}
	categoryCacheDao := &cacheDaos.AllCategoryIdsRedisDao{Pool: redisCachePool}
	categoryToProductCacheDao := &cacheDaos.CategoryIdToProductIdRedisDao{Pool: redisCachePool}

	// Get db results
	apiDbResult, dbReadErr := readFromDb(ctx)
	if dbReadErr != nil {
		ctx.JSON(http.StatusInternalServerError, Home_Db_Read_Err_Code)
		return
	}
	productIdMap := apiCommons.NewProductIdMap(apiDbResult.ProductsMap)

	// Make all sections
	newItemsProductRailSection := section.NewItemsProductRailSection(ctx, productCacheDao, productIdMap, apiDbResult)
	randomProductFullSection := section.RandomProductFullSection(ctx, productCacheDao, productIdMap, apiDbResult)
	categoryRailSection := section.CategoryRailSection(ctx, categoryCacheDao, apiDbResult)
	categoryToProductRailSections := section.CategoryToProductRailSections(ctx, categoryToProductCacheDao, productIdMap, apiDbResult)
	randomProduct2FullSection := section.RandomProductFullSection(ctx, productCacheDao, productIdMap, apiDbResult)
	remainingDualProductSection := section.RemainingProductsSection(productIdMap.RemainingProducts(), apiDbResult)

	// Arrange sections
	var snippetSectionDataList []*snippets.SnippetSectionData
	snippetSectionDataList = apiCommons.AppendIgnoringNils(snippetSectionDataList, newItemsProductRailSection)
	snippetSectionDataList = apiCommons.AppendIgnoringNils(snippetSectionDataList, randomProductFullSection)
	snippetSectionDataList = apiCommons.AppendIgnoringNils(snippetSectionDataList, categoryRailSection)
	snippetSectionDataList = apiCommons.AppendIgnoringNils(snippetSectionDataList, categoryToProductRailSections...)
	snippetSectionDataList = apiCommons.AppendIgnoringNils(snippetSectionDataList, randomProduct2FullSection)
	snippetSectionDataList = apiCommons.AppendIgnoringNils(snippetSectionDataList, remainingDualProductSection)

	homeApiResponse := models.HomeApiResponse{
		Status:   "success",
		Message:  "",
		Snippets: snippetSectionDataList,
	}
	ctx.JSON(http.StatusOK, homeApiResponse)
}
