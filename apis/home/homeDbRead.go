package home

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/hashicorp/go-multierror"
	"phantom/apis/apiCommons"
	"phantom/apis/home/models"
	"phantom/dataLayer/databasDaos"
	"phantom/dataLayer/dbModels"
	"phantom/ginRouter"
	"sync"
)

func readFromDb(ctx *gin.Context) (models.ApiDbResult, error) {
	db := ctx.MustGet(ginRouter.SQL_DB).(*sql.DB)

	// Define our results
	parallelization := 3
	productsMap := map[int64]*dbModels.Product{}
	categoriesMap := map[int64]*dbModels.Category{}
	brandsMap := map[int64]*dbModels.Brand{}

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
		productsMap = readProductsFromDb(ctx, db, errorChan, wg)
	}()
	go func() {
		categoriesMap = readCategoriesFromDb(ctx, db, errorChan, wg)
	}()
	go func() {
		brandsMap = readBrandsFromDb(ctx, db, errorChan, wg)
	}()

	// Combine all errors we have and handle if needed
	var combinedErr error
	for err := range errorChan {
		combinedErr = multierror.Append(combinedErr, err)
	}
	if combinedErr != nil {
		return models.EmptyHomeApiDbResult(), combinedErr
	}

	// Return actual db results
	return models.ApiDbResult{
		ProductsMap:   productsMap,
		CategoriesMap: categoriesMap,
		BrandsMap:     brandsMap,
	}, nil
}

func readProductsFromDb(
	ctx *gin.Context,
	db *sql.DB,
	errChan chan error,
	wg *sync.WaitGroup,
) map[int64]*dbModels.Product {
	defer wg.Done()
	productDbDao := databasDaos.ProductSqlDao{DB: db}
	products, err := productDbDao.ReadAllProducts()
	if err != nil {
		logData := apiCommons.NewApiErrorLogData(ctx, "Error reading all products from db", err)
		apiCommons.LogApiError(logData)
		errChan <- err
	} else {
		productMap := make(map[int64]*dbModels.Product, 0)
		for index := range *products {
			product := (*products)[index]
			productMap[product.Id] = &product
		}
		return productMap
	}
	return nil
}

func readCategoriesFromDb(
	ctx *gin.Context,
	db *sql.DB,
	errChan chan error,
	wg *sync.WaitGroup,
) map[int64]*dbModels.Category {
	defer wg.Done()
	categoryDbDao := databasDaos.CategorySqlDao{DB: db}
	categories, err := categoryDbDao.ReadAllCategories()
	if err != nil {
		logData := apiCommons.NewApiErrorLogData(ctx, "Error reading all categories from db", err)
		apiCommons.LogApiError(logData)
		errChan <- err
	} else {
		categoryMap := make(map[int64]*dbModels.Category, 0)
		for index := range *categories {
			category := (*categories)[index]
			categoryMap[category.Id] = &category
		}
		return categoryMap
	}
	return nil
}

func readBrandsFromDb(
	ctx *gin.Context,
	db *sql.DB,
	errChan chan error,
	wg *sync.WaitGroup,
) map[int64]*dbModels.Brand {
	defer wg.Done()
	brandDbDao := databasDaos.BrandSqlDao{DB: db}
	brands, err := brandDbDao.ReadAllBrands()
	if err != nil {
		logData := apiCommons.NewApiErrorLogData(ctx, "Error reading all brands from db", err)
		apiCommons.LogApiError(logData)
		errChan <- err
	} else {
		brandMap := make(map[int64]*dbModels.Brand, 0)
		for index := range *brands {
			brand := (*brands)[index]
			brandMap[brand.Id] = &brand
		}
		return brandMap
	}
	return nil
}
