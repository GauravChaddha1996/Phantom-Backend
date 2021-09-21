package filter

import (
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	"net/http"
	"phantom/apis/apiCommons"
	"phantom/apis/filter/models"
	"phantom/dataLayer/cacheDaos"
	"phantom/dataLayer/databaseDaos"
	"phantom/dataLayer/dbModels"
	"phantom/ginRouter"
	"sort"
)

func ApiHandler(ctx *gin.Context) {
	// Initialize or find dependencies
	redisPool := ctx.MustGet(ginRouter.REDIS_POOL).(*redis.Pool)
	sqlDb := ctx.MustGet(ginRouter.SQL_DB).(*sql.DB)
	filterProductsDao := cacheDaos.FilterProductsDao{Pool: redisPool}
	productDb := databaseDaos.ProductSqlDao{DB: sqlDb}

	// Read api request model
	apiRequestModel, apiRequestReadErr := models.ReadApiRequestModel(ctx)
	if apiRequestReadErr != nil {
		ctx.JSON(http.StatusInternalServerError, apiRequestReadErr.Error())
		return
	}

	// Find filtered product ids
	productIds, filterDaoReadErr := filterProductsDao.FindProductsForFilter(apiRequestModel.CategoryId, apiRequestModel.PropertyValueIds)
	if filterDaoReadErr != nil {
		msg := "something went wrong while filtering products"
		apiCommons.LogApiError(apiCommons.NewApiErrorLogData(ctx, msg, filterDaoReadErr))
		ctx.JSON(http.StatusInternalServerError, errors.New(msg))
		return
	}

	// Read filtered products data from db
	products, dbReadErr := productDb.ReadProducts(*productIds)
	if dbReadErr != nil {
		msg := "something went wrong while reading from db"
		apiCommons.LogApiError(apiCommons.NewApiErrorLogData(ctx, msg, dbReadErr))
		ctx.JSON(http.StatusInternalServerError, errors.New(msg))
		return
	}

	// Sort on api request model basis
	sortProducts(products)

	// Make api response
	ctx.JSON(http.StatusOK, products)
}

func sortProducts(products *[]dbModels.Product) {
	sort.SliceStable(*products, func(first, second int) bool {
		productFirst := (*products)[first]
		productSecond := (*products)[second]
		return productFirst.CreatedAt.After(*productSecond.CreatedAt)
	})
}
