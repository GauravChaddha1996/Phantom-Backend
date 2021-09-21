package filter

import (
	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	"net/http"
	"phantom/apis/apiCommons"
	"phantom/apis/filter/models"
	"phantom/apis/filter/section"
	"phantom/dataLayer/cacheDaos"
	"phantom/dataLayer/uiModels/snippets"
	"phantom/ginRouter"
)

const Filter_Api_Request_Model_Read_Err = "Err code: 1"
const Filter_Api_Filtering_Err = "Err code: 2"
const Filter_Api_Db_Read_Err = "Err code: 3"
const Filter_Api_Property_For_Category_Err = "Err code: 4"

func ApiHandler(ctx *gin.Context) {
	// Initialization or find dependencies
	redisPool := ctx.MustGet(ginRouter.REDIS_POOL).(*redis.Pool)

	// Read api request model
	apiRequest, apiRequestReadErr := models.ReadApiRequestModel(ctx)
	if apiRequestReadErr != nil {
		ctx.JSON(http.StatusInternalServerError, Filter_Api_Request_Model_Read_Err)
		return
	}

	// Find filtered product ids
	productIds, filteringErr := findFilteredProductIds(ctx, redisPool, apiRequest)
	if filteringErr != nil {
		ctx.JSON(http.StatusInternalServerError, Filter_Api_Filtering_Err)
		return
	}

	// Find property ids of category id
	propertyIds, propertyForCategoryErr := findPropertyIdsOfCategoryId(ctx, redisPool, apiRequest.CategoryId)
	if propertyForCategoryErr != nil {
		ctx.JSON(http.StatusInternalServerError, Filter_Api_Property_For_Category_Err)
		return
	}

	// Read results needed from db
	apiDbResult, dbReadErr := readFromDb(ctx, productIds, propertyIds)
	if dbReadErr != nil {
		ctx.JSON(http.StatusInternalServerError, Filter_Api_Db_Read_Err)
		return
	}

	// Sort on api request model basis
	//sortProducts(apiDbResult, apiRequest)

	// Make api response
	response := makeFilterApiResponse(apiRequest, apiDbResult)
	ctx.JSON(http.StatusOK, response)
}

func findFilteredProductIds(
	ctx *gin.Context,
	redisPool *redis.Pool,
	apiRequest *models.ApiRequest,
) (*[]int64, error) {
	filterProductsDao := cacheDaos.FilterProductsDao{Pool: redisPool}
	productIds, err := filterProductsDao.FindProductsForFilter(apiRequest.CategoryId, apiRequest.PropertyValueIds)
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

//func sortProducts(apiDbResult *models.ApiDbResult, apiRequest *models.ApiRequest) {
//	sort.SliceStable(*apiDbResult, func(one, two int) bool {
//		productOne := (*apiDbResult)[one]
//		productTwo := (*apiDbResult)[two]
//
//		isProductOneBeforeProductTwo := true
//
//		for _, sortMethod := range models.AllSortMethods {
//			switch apiRequest.SortId {
//			case models.Sort_Newly_Added:
//				isProductOneBeforeProductTwo = productOne.CreatedAt.After(*productTwo.CreatedAt)
//			case models.Sort_Price_Low_To_High:
//				isProductOneBeforeProductTwo = productOne.Cost < productTwo.Cost
//			case models.Sort_Price_High_To_Low:
//				isProductOneBeforeProductTwo = productOne.Cost > productTwo.Cost
//			case models.Sort_Brand_Wise:
//				isProductOneBeforeProductTwo = false
//			case models.Sort_A_Z:
//				isProductOneBeforeProductTwo = strings.Compare(productOne.Name, productTwo.Name) < 0
//			}
//		}
//		return isProductOneBeforeProductTwo
//	})
//}

func makeFilterApiResponse(apiRequest *models.ApiRequest, apiDbResult *models.ApiDbResult) models.FilterApiResponse {
	snippetSectionData := section.GetFilteredProductSnippetSection(apiDbResult)
	return models.FilterApiResponse{
		Status:            "success",
		Message:           "",
		Snippets:          []*snippets.SnippetSectionData{&snippetSectionData},
		FilterSheetUiData: models.MakeFilterSheetUiData(apiRequest, apiDbResult.PropertyToPropertyValueMap),
		SortSheetUiData:   models.MakeSortSheetUiData(apiRequest.SortId),
	}
}
