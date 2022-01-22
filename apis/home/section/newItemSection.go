package section

import (
	"github.com/gin-gonic/gin"
	"phantom/apis/apiCommons"
	"phantom/apis/home/models"
	"phantom/dataLayer/cacheDaos"
	"phantom/dataLayer/uiModels/atoms"
	"phantom/dataLayer/uiModels/snippets"
)

const NewlyIntroducedProductIdsTotalCount = 8
const newlyIntroducedSectionItemCount = 4
const newlyIntroducedSectionHeader = "Newly introduced"

func NewItemsProductRailSection(
	ctx *gin.Context,
	productCacheDao *cacheDaos.AllProductIdsRedisDao,
	productIdMap *apiCommons.ProductIdMap,
	apiDbResult models.ApiDbResult,
) *snippets.SnippetSectionData {
	var productRailSnippets []snippets.ProductRailSnippetData

	newProductIds, err := productCacheDao.ReadFirstNProductIds(NewlyIntroducedProductIdsTotalCount)
	if err != nil {
		logData := apiCommons.NewApiErrorLogData(ctx, "Error reading first n products from cache", err)
		apiCommons.LogApiError(logData)
		return nil
	}
	apiDbResult.NewProductIdsMap.PutAllInt64s(newProductIds)

	newProductIdsForSection := (*newProductIds)[:newlyIntroducedSectionItemCount]
	productIdMap.PutAllInt64s(&newProductIdsForSection)
	for _, productId := range newProductIdsForSection {
		product := apiDbResult.ProductsMap[productId]
		category := apiDbResult.CategoriesMap[product.CategoryId]
		brand := apiDbResult.BrandsMap[product.BrandId]
		snippet := snippets.MakeProductRailSnippet(*product, *category, *brand, true)
		productRailSnippets = append(productRailSnippets, snippet)
	}

	return &snippets.SnippetSectionData{
		HeaderData: &snippets.SnippetSectionHeaderData{
			Title: &atoms.TextData{
				Text: newlyIntroducedSectionHeader,
			},
		},
		Snippets: &productRailSnippets,
	}
}
