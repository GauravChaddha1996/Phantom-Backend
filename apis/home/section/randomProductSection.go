package section

import (
	"github.com/gin-gonic/gin"
	"phantom/apis/apiCommons"
	"phantom/apis/home/models"
	"phantom/dataLayer/cacheDaos"
	"phantom/dataLayer/uiModels/snippets"
)

const randomProductSectionMaxIteration = 20

func RandomProductFullSection(
	ctx *gin.Context,
	productCacheDao *cacheDaos.AllProductIdsRedisDao,
	productIdMap *apiCommons.ProductIdMap,
	apiDbResult models.ApiDbResult,
) *snippets.SnippetSectionData {
	var productFullSnippets []snippets.ProductFullSnippet

	randomProductId := findRandomProductId(ctx, productCacheDao, productIdMap)
	if randomProductId == nil {
		logData := apiCommons.NewApiErrorLogData(ctx, "Error finding random product id", nil)
		apiCommons.LogApiError(logData)
		return nil
	}
	
	product := apiDbResult.ProductsMap[*randomProductId]
	category := apiDbResult.CategoriesMap[product.CategoryId]
	brand := apiDbResult.BrandsMap[product.BrandId]
	snippet := snippets.MakeProductFullSnippet(*product, *category, *brand)
	productFullSnippets = append(productFullSnippets, snippet)

	return &snippets.SnippetSectionData{
		Type:     snippets.ProductFullSection,
		Snippets: apiCommons.ToBaseSnippets(productFullSnippets),
	}
}

func findRandomProductId(
	ctx *gin.Context,
	productCacheDao *cacheDaos.AllProductIdsRedisDao,
	productIdMap *apiCommons.ProductIdMap,
) *int64 {
	isRandomProductIdNotFound := true
	totalIterations := 0

	for isRandomProductIdNotFound && totalIterations < randomProductSectionMaxIteration {
		totalIterations++
		randomProductId, err := productCacheDao.ReadRandomProduct()
		if err == nil {
			if productIdMap.Contains(*randomProductId) == false {
				isRandomProductIdNotFound = false
				productIdMap.Put(*randomProductId)
				return randomProductId
			}
		} else {
			logData := apiCommons.NewApiErrorLogData(ctx, "Error reading random product from cache", err)
			apiCommons.LogApiError(logData)
			isRandomProductIdNotFound = false
		}
	}
	return nil
}
