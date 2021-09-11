package section

import (
	"phantom/apis/apiCommons"
	"phantom/apis/home/models"
	"phantom/dataLayer/cacheDaos"
	"phantom/dataLayer/uiModels/atoms"
	"phantom/dataLayer/uiModels/snippets"
)

const newlyIntroducedSectionItemCount = 2
const newlyIntroducedSectionHeader = "Newly introduced"

func NewItemsProductRailSection(
	productCacheDao *cacheDaos.AllProductIdsRedisDao,
	productIdMap *apiCommons.ProductIdMap,
	apiDbResult models.ApiDbResult,
) snippets.SnippetSectionData {
	var productRailSnippets []snippets.ProductRailSnippet

	productIds, err := productCacheDao.ReadFirstNProductIds(newlyIntroducedSectionItemCount)
	if err == nil {
		productIdMap.PutAllInt64s(productIds)
		for _, productId := range *productIds {
			snippet := getProductRailSnippet(productId, apiDbResult)
			productRailSnippets = append(productRailSnippets, snippet)
		}
	}

	return snippets.SnippetSectionData{
		Type: snippets.ProductRailSection,
		HeaderData: &snippets.SnippetSectionHeaderData{
			Title: &atoms.TextData{Text: newlyIntroducedSectionHeader},
		},
		Snippets: apiCommons.ToBaseSnippets(productRailSnippets),
	}
}
