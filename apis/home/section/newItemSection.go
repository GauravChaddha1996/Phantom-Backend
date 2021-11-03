package section

import (
	"github.com/gin-gonic/gin"
	"phantom/apis/apiCommons"
	"phantom/apis/home/models"
	"phantom/dataLayer/cacheDaos"
	"phantom/dataLayer/uiModels/atoms"
	"phantom/dataLayer/uiModels/snippets"
)

const newlyIntroducedSectionItemCount = 2
const newlyIntroducedSectionHeader = "Newly introduced"

func NewItemsProductRailSection(
	ctx *gin.Context,
	productCacheDao *cacheDaos.AllProductIdsRedisDao,
	productIdMap *apiCommons.ProductIdMap,
	apiDbResult models.ApiDbResult,
) *snippets.SnippetSectionData {
	var productRailSnippets []snippets.ProductRailSnippetData

	productIds, err := productCacheDao.ReadFirstNProductIds(newlyIntroducedSectionItemCount)
	if err != nil {
		logData := apiCommons.NewApiErrorLogData(ctx, "Error reading first n products from cache", err)
		apiCommons.LogApiError(logData)
		return nil
	}

	productIdMap.PutAllInt64s(productIds)
	for _, productId := range *productIds {
		product := apiDbResult.ProductsMap[productId]
		category := apiDbResult.CategoriesMap[product.CategoryId]
		brand := apiDbResult.BrandsMap[product.BrandId]
		snippet := snippets.MakeProductRailSnippet(*product, *category, *brand)
		productRailSnippets = append(productRailSnippets, snippet)
	}

	return &snippets.SnippetSectionData{
		Type: snippets.ProductRailSnippet,
		HeaderData: &snippets.SnippetSectionHeaderData{
			Title: &atoms.TextData{Text: newlyIntroducedSectionHeader},
		},
		Snippets: &productRailSnippets,
	}
}
