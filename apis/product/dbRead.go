package product

import (
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/hashicorp/go-multierror"
	"go.uber.org/multierr"
	"phantom/apis/apiCommons"
	"phantom/apis/product/models"
	"phantom/dataLayer/databaseDaos"
	"phantom/dataLayer/dbModels"
	"phantom/ginRouter"
	"sync"
)

func readFromDb(ctx *gin.Context, apiRequest *models.ApiRequest) (*models.ApiDbResult, error) {
	// Initialization or find dependencies
	db := ctx.MustGet(ginRouter.SQL_DB).(*sql.DB)
	productDao := databaseDaos.ProductSqlDao{DB: db}

	// Define our results
	parallelization := 4
	productId := apiRequest.ProductId
	var product *dbModels.Product
	var productImages []dbModels.ProductImage
	var brand *dbModels.Brand
	var category *dbModels.Category
	var propertyMapping *map[dbModels.Property][]dbModels.PropertyValue

	// Setup for running all db tasks parallelly and waiting for them
	wg := &sync.WaitGroup{}
	wg.Add(parallelization)
	errorChan := make(chan error)
	go func() {
		wg.Wait()
		close(errorChan)
	}()

	// Do the first serialized task need for all parallelization task
	product, productReadErr := productDao.ReadProduct(productId)
	if productReadErr != nil {
		msg := "error reading product from db"
		logData := apiCommons.NewApiErrorLogData(ctx, msg, productReadErr)
		apiCommons.LogApiError(logData)
		return nil, errors.New(msg)
	}

	// Run all db tasks parallelly
	go func() {
		productImages = readProductImages(ctx, db, productId, errorChan, wg)
	}()
	go func() {
		brand = readBrand(ctx, db, product.BrandId, errorChan, wg)
	}()
	go func() {
		category = readCategory(ctx, db, product.CategoryId, errorChan, wg)
	}()
	go func() {
		propertyMapping = readPropertyMapping(ctx, db, productId, errorChan, wg)
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
		Product:         product,
		ProductImages:   productImages,
		Category:        category,
		Brand:           brand,
		PropertyMapping: propertyMapping,
	}, nil
}

func readProductImages(
	ctx *gin.Context, db *sql.DB, productId int64, errChan chan error, wg *sync.WaitGroup,
) []dbModels.ProductImage {
	defer wg.Done()
	productImagesDao := databaseDaos.ProductImageSqlDao{DB: db}
	productImages, err := productImagesDao.ReadProductImages(productId)
	if err != nil {
		msg := "error reading product images from db"
		logData := apiCommons.NewApiErrorLogData(ctx, msg, err)
		apiCommons.LogApiError(logData)
		errChan <- errors.New(msg)
	} else {
		return *productImages
	}
	return nil
}

func readBrand(ctx *gin.Context, db *sql.DB, brandId int64, errChan chan error, wg *sync.WaitGroup) *dbModels.Brand {
	defer wg.Done()
	brandDao := databaseDaos.BrandSqlDao{DB: db}
	brand, err := brandDao.ReadBrandComplete(brandId)
	if err != nil {
		msg := "error reading brand from db"
		logData := apiCommons.NewApiErrorLogData(ctx, msg, err)
		apiCommons.LogApiError(logData)
		errChan <- errors.New(msg)
	} else {
		return brand
	}
	return nil
}

func readCategory(
	ctx *gin.Context, db *sql.DB, categoryId int64, errChan chan error, wg *sync.WaitGroup,
) *dbModels.Category {
	defer wg.Done()
	categoryDao := databaseDaos.CategorySqlDao{DB: db}
	category, err := categoryDao.ReadCategoryComplete(categoryId)
	if err != nil {
		msg := "error reading category from db"
		logData := apiCommons.NewApiErrorLogData(ctx, msg, err)
		apiCommons.LogApiError(logData)
		errChan <- errors.New(msg)
	} else {
		return category
	}
	return nil
}

func readPropertyMapping(
	ctx *gin.Context, db *sql.DB, productId int64, errChan chan error, wg *sync.WaitGroup,
) *map[dbModels.Property][]dbModels.PropertyValue {
	defer wg.Done()
	productToPropertyDao := databaseDaos.ProductToPropertySqlDao{DB: db}
	propertyDao := databaseDaos.PropertySqlDao{DB: db}
	propertyValueDao := databaseDaos.PropertyValueSqlDao{DB: db}

	propertyMapping, propertyMappingErr := productToPropertyDao.ReadAllProductToPropertyMappingForProductId(productId)
	if propertyMappingErr != nil {
		msg := "error reading property mapping from db"
		logData := apiCommons.NewApiErrorLogData(ctx, msg, propertyMappingErr)
		apiCommons.LogApiError(logData)
		errChan <- errors.New(msg)
		return nil
	}

	propertyMap := make(map[dbModels.Property][]dbModels.PropertyValue)
	var propertyMappingItemReadErrArr []error
	for _, productToProperty := range *propertyMapping {
		property, propertyReadErr := propertyDao.ReadPropertyComplete(productToProperty.PropertyId)

		if propertyReadErr != nil {
			propertyMappingItemReadErrArr = append(propertyMappingItemReadErrArr, propertyReadErr)
			continue
		}

		propertyValue, propertyValueReadErr := propertyValueDao.ReadPropertyValueComplete(productToProperty.ValueId)
		if propertyValueReadErr != nil {
			propertyMappingItemReadErrArr = append(propertyMappingItemReadErrArr, propertyValueReadErr)
			continue
		}

		currentArr := propertyMap[*property]
		if currentArr == nil {
			currentArr = make([]dbModels.PropertyValue, 0)
		}
		currentArr = append(currentArr, *propertyValue)
		propertyMap[*property] = currentArr
	}

	combinedPropertyMappingReadErr := multierr.Combine(propertyMappingItemReadErrArr...)
	if combinedPropertyMappingReadErr != nil {
		msg := "error reading property mapping item(s) from db"
		logData := apiCommons.NewApiErrorLogData(ctx, msg, combinedPropertyMappingReadErr)
		apiCommons.LogApiError(logData)
		errChan <- errors.New(msg)
	} else {
		return &propertyMap
	}
	return nil
}
