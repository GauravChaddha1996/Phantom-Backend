package section

import (
	"phantom/apis/apiCommons"
	"phantom/apis/home/models"
	"phantom/dataLayer/cacheDaos"
	"phantom/dataLayer/uiModels/snippets"
)

const randomProductSectionMaxIteration = 20

func RandomProductFullSection(
	productCacheDao *cacheDaos.AllProductIdsRedisDao,
	productIdMap *apiCommons.ProductIdMap,
	apiDbResult models.ApiDbResult,
) snippets.SnippetSectionData {
	var productFullSnippets []snippets.ProductFullSnippet

	randomProductId := findRandomProductId(productCacheDao, productIdMap)
	if randomProductId != nil {
		snippet := getProductFullSnippet(*randomProductId, apiDbResult)
		productFullSnippets = append(productFullSnippets, snippet)
	}

	return snippets.SnippetSectionData{
		Type:     snippets.ProductFullSection,
		Snippets: apiCommons.ToBaseSnippets(productFullSnippets),
	}
}

func findRandomProductId(
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
			isRandomProductIdNotFound = false
		}
	}
	return nil
}
