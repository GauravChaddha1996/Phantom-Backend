package filter

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	"net/http"
	"phantom/apis/apiCommons"
	"phantom/apis/filter/models"
	"phantom/apis/filter/section"
	"phantom/dataLayer"
	"phantom/dataLayer/cacheDaos"
	"phantom/dataLayer/databaseDaos"
	"phantom/dataLayer/dbModels"
	"phantom/dataLayer/uiModels/atoms"
	"phantom/dataLayer/uiModels/snippets"
	"phantom/ginRouter"
	"sort"
)

const apiFilteringErr = "Err code: 1"
const apiDbReadErr = "Err code: 2"
const apiPropertyForCategoryErr = "Err code: 3"
const apiCategoryReadErr = "Err code: 4"
const apiProductCacheReadErr = "Err code: 5"
const newlyIntroducedProductIdsTotalCount = 8

func ApiHandler(ctx *gin.Context) {
	// Initialization or find dependencies
	redisCachePool := ctx.MustGet(ginRouter.REDIS_POOL).(*redis.Pool)
	db := ctx.MustGet(ginRouter.SQL_DB).(*sql.DB)
	categoryDbDao := databaseDaos.CategorySqlDao{DB: db}
	productCacheDao := cacheDaos.AllProductIdsRedisDao{Pool: redisCachePool}

	// Read api request model
	apiRequest, apiRequestReadErr := models.ReadApiRequestModel(ctx)
	if apiRequestReadErr != nil {
		ctx.JSON(http.StatusInternalServerError, apiRequestReadErr.Error())
		return
	}

	// Find category data
	category, categoryReadErr := categoryDbDao.ReadCategoryComplete(apiRequest.CategoryId)
	if categoryReadErr != nil {
		ctx.JSON(http.StatusInternalServerError, apiCategoryReadErr)
		return
	}

	// Find filtered product ids
	productIds, filteringErr := findFilteredProductIds(ctx, redisCachePool, apiRequest)
	if filteringErr != nil {
		ctx.JSON(http.StatusInternalServerError, apiFilteringErr)
		return
	}

	// Find property ids of category id
	propertyIds, propertyForCategoryErr := findPropertyIdsOfCategoryId(ctx, redisCachePool, apiRequest.CategoryId)
	if propertyForCategoryErr != nil {
		ctx.JSON(http.StatusInternalServerError, apiPropertyForCategoryErr)
		return
	}

	// Read results needed from db
	apiDbResult, dbReadErr := readFromDb(ctx, productIds, propertyIds)
	if dbReadErr != nil {
		ctx.JSON(http.StatusInternalServerError, apiDbReadErr)
		return
	}

	// Set newly introduced product ids in api db result
	newProductIds, err := productCacheDao.ReadFirstNProductIds(newlyIntroducedProductIdsTotalCount)
	if err != nil {
		logData := apiCommons.NewApiErrorLogData(ctx, "Error reading first n products from cache", err)
		apiCommons.LogApiError(logData)
		ctx.JSON(http.StatusInternalServerError, apiProductCacheReadErr)
		return
	}
	apiDbResult.NewProductIdsMap.PutAllInt64s(newProductIds)

	// Sort on api request model basis
	sortProducts(apiDbResult, apiRequest)

	// Make api response
	response := makeFilterApiResponse(apiRequest, apiDbResult, category)
	ctx.JSON(http.StatusOK, response)
}

func findFilteredProductIds(
	ctx *gin.Context,
	redisPool *redis.Pool,
	apiRequest *models.ApiRequest,
) (*[]int64, error) {
	filterProductsDao := cacheDaos.FilterProductsDao{Pool: redisPool}
	productIds, err := filterProductsDao.FindProductsForFilter(apiRequest.CategoryId, apiRequest.PropertyValueIdsMap)
	if err != nil {
		logData := apiCommons.NewApiErrorLogData(ctx, "something went wrong while filtering products", err)
		apiCommons.LogApiError(logData)
		return nil, err
	}
	return productIds, nil
}

func findPropertyIdsOfCategoryId(
	ctx *gin.Context,
	redisPool *redis.Pool,
	categoryId int64,
) (*[]int64, error) {
	categoryIdToPropertyIdRedisDao := cacheDaos.CategoryIdToPropertyIdRedisDao{Pool: redisPool}
	propertyIds, err := categoryIdToPropertyIdRedisDao.ReadPropertyIdsForCategoryId(categoryId)
	if err != nil {
		logData := apiCommons.NewApiErrorLogData(ctx, "something went wrong while finding property ids of category id", err)
		apiCommons.LogApiError(logData)
		return nil, err
	}
	return propertyIds, nil
}

func sortProducts(apiDbResult *models.ApiDbResult, apiRequest *models.ApiRequest) {
	productList := apiDbResult.ProductsList
	sortMethod := dataLayer.FindSortMethod(apiRequest.SortId)
	sort.SliceStable(productList, func(one, two int) bool {
		productOne := (productList)[one]
		productTwo := (productList)[two]
		return sortMethod.Compare(productOne, productTwo)
	})
}

func makeFilterApiResponse(
	apiRequest *models.ApiRequest, apiDbResult *models.ApiDbResult, category *dbModels.Category,
) models.FilterApiResponse {
	snippetSectionData, sectionHeader := section.GetFilteredProductSnippetSection(apiDbResult)
	return models.FilterApiResponse{
		Status:    "success",
		Message:   "",
		PageTitle: atoms.TextData{Text: category.Name},
		SnippetSectionHeader: atoms.TextData{
			Text: sectionHeader,
			Font: &atoms.FontData{Style: atoms.FontTitleMedium},
		},
		SnippetSectionList: []*snippets.SnippetSectionData{&snippetSectionData},
		FilterSheetUiData:  models.MakeFilterSheetUiData(apiRequest, apiDbResult.PropertyToPropertyValueMap),
		SortSheetUiData:    models.MakeSortSheetUiData(apiRequest.SortId),
	}
}
