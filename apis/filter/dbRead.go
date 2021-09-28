package filter

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hashicorp/go-multierror"
	"github.com/spf13/cast"
	"phantom/apis/apiCommons"
	"phantom/apis/filter/models"
	"phantom/dataLayer/databaseDaos"
	"phantom/dataLayer/dbModels"
	"phantom/ginRouter"
	"sync"
)

func readFromDb(ctx *gin.Context, productIds *[]int64, propertyIds *[]int64) (*models.ApiDbResult, error) {
	// Initialization or find dependencies
	sqlDb := ctx.MustGet(ginRouter.SQL_DB).(*sql.DB)

	// Define our results
	parallelization := 2
	var productsList []dbModels.Product
	brandsMap := map[int64]dbModels.Brand{}
	propertyToPropertyValueMap := map[dbModels.Property][]dbModels.PropertyValue{}

	// Setup for running all db tasks parallelly and waiting for them
	wg := &sync.WaitGroup{}
	wg.Add(parallelization)
	errorChan := make(chan error)
	go func() {
		wg.Wait()
		close(errorChan)
	}()

	// Run all db tasks parallelly
	go func() {
		productsList = readProductsFromDb(ctx, sqlDb, productIds, errorChan, wg)
		var brandIdsToFind []int64
		for _, product := range productsList {
			brandIdsToFind = append(brandIdsToFind, product.BrandId)
		}
		brandsMap = readBrandsFromDb(ctx, sqlDb, errorChan, brandIdsToFind)
	}()
	go func() {
		propertyToPropertyValueMap = readPropertyAndPropertyValueFromDb(ctx, sqlDb, errorChan, wg, propertyIds)
	}()

	// Combine all errors we have and handle if needed
	var combinedErr error
	for err := range errorChan {
		combinedErr = multierror.Append(combinedErr, err)
	}
	if combinedErr != nil {
		return nil, combinedErr
	}

	// Return actual db results
	return &models.ApiDbResult{
		ProductsList:               productsList,
		BrandsMap:                  brandsMap,
		PropertyToPropertyValueMap: propertyToPropertyValueMap,
	}, nil
}

func readProductsFromDb(
	ctx *gin.Context, db *sql.DB, productIds *[]int64, errChan chan error, wg *sync.WaitGroup,
) []dbModels.Product {
	defer wg.Done()
	productDbDao := databaseDaos.ProductSqlDao{DB: db}
	products, err := productDbDao.ReadProducts(*productIds)
	if err != nil {
		logData := apiCommons.NewApiErrorLogData(ctx, "Error reading products from db", err)
		apiCommons.LogApiError(logData)
		errChan <- err
	} else {
		return *products
	}
	return nil
}

func readBrandsFromDb(
	ctx *gin.Context,
	db *sql.DB,
	errChan chan error,
	brandIds []int64,
) map[int64]dbModels.Brand {
	brandDbDao := databaseDaos.BrandSqlDao{DB: db}
	brands, err := brandDbDao.ReadBrands(brandIds)
	if err != nil {
		logData := apiCommons.NewApiErrorLogData(ctx, "Error reading brands from db", err)
		apiCommons.LogApiError(logData)
		errChan <- err
	} else {
		brandMap := make(map[int64]dbModels.Brand, 0)
		for index := range *brands {
			brand := (*brands)[index]
			brandMap[brand.Id] = brand
		}
		return brandMap
	}
	return nil
}

func readPropertyAndPropertyValueFromDb(
	ctx *gin.Context,
	db *sql.DB,
	errChan chan error,
	wg *sync.WaitGroup,
	propertyIds *[]int64,
) map[dbModels.Property][]dbModels.PropertyValue {
	defer wg.Done()
	propertyDbDao := databaseDaos.PropertySqlDao{DB: db}
	propertyValueDbDao := databaseDaos.PropertyValueSqlDao{DB: db}

	properties, propertiesReadErr := propertyDbDao.ReadProperties(*propertyIds)
	if propertiesReadErr != nil {
		logData := apiCommons.NewApiErrorLogData(ctx, "Error reading properties", propertiesReadErr)
		logData.Data["property_ids"] = fmt.Sprint(*propertyIds)
		apiCommons.LogApiError(logData)
		errChan <- propertiesReadErr
	} else {
		return makePropertyToPropertyValueMap(ctx, properties, propertyValueDbDao, errChan, propertiesReadErr)
	}
	return nil
}

func makePropertyToPropertyValueMap(
	ctx *gin.Context,
	properties *[]dbModels.Property,
	propertyValueDbDao databaseDaos.PropertyValueSqlDao,
	errChan chan error,
	propertiesReadErr error,
) map[dbModels.Property][]dbModels.PropertyValue {
	propertyToPropertyValueMap := make(map[dbModels.Property][]dbModels.PropertyValue, 0)
	for _, property := range *properties {
		propertyValuesOfProperty, propertyValueReadErr := propertyValueDbDao.ReadAllPropertyValuesOfProperty(property.Id)
		if propertyValueReadErr != nil {
			logData := apiCommons.NewApiErrorLogData(ctx, "Error reading property values of property", propertyValueReadErr)
			logData.Data["property_id"] = cast.ToString(property.Id)
			apiCommons.LogApiError(logData)
			errChan <- propertiesReadErr
		} else {
			propertyToPropertyValueMap[property] = *propertyValuesOfProperty
		}
	}
	return propertyToPropertyValueMap
}
