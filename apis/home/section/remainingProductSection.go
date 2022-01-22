package section

import (
	"phantom/apis/home/models"
	"phantom/dataLayer/uiModels/atoms"
	"phantom/dataLayer/uiModels/snippets"
)

const remainingProductSectionHeader = "Exciting products!"

func RemainingProductsSection(
	remainingProductIds *[]int64,
	apiDbResult models.ApiDbResult,
) *snippets.SnippetSectionData {

	var productDualSnippets []snippets.ProductDualSnippetData
	for _, productId := range *remainingProductIds {
		product := apiDbResult.ProductsMap[productId]
		brand := apiDbResult.BrandsMap[product.BrandId]
		snippet := snippets.MakeProductDualSnippet(*product, *brand, apiDbResult.NewProductIdsMap.Contains(productId))
		productDualSnippets = append(productDualSnippets, snippet)
	}

	return &snippets.SnippetSectionData{
		HeaderData: &snippets.SnippetSectionHeaderData{
			Title: &atoms.TextData{
				Text: remainingProductSectionHeader,
			},
			Subtitle:    nil,
			RightButton: nil,
		},
		Snippets: &productDualSnippets,
	}
}
