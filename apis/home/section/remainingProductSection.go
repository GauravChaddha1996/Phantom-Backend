package section

import (
	"phantom/apis/apiCommons"
	"phantom/apis/home/models"
	"phantom/dataLayer/uiModels/atoms"
	"phantom/dataLayer/uiModels/snippets"
)

const remainingProductSectionHeader = "Exciting products!"

func RemainingProductsSection(
	remainingProductIds *[]int64,
	apiDbResult models.ApiDbResult,
) *snippets.SnippetSectionData {

	var productDualSnippets []snippets.ProductDualSnippet
	for _, productId := range *remainingProductIds {
		snippet := getProductDualSnippet(productId, apiDbResult)
		productDualSnippets = append(productDualSnippets, snippet)
	}

	return &snippets.SnippetSectionData{
		Type: snippets.ProductDualSection,
		HeaderData: &snippets.SnippetSectionHeaderData{
			Title:       &atoms.TextData{Text: remainingProductSectionHeader},
			Subtitle:    nil,
			RightButton: nil,
		},
		Snippets: apiCommons.ToBaseSnippets(productDualSnippets),
	}
}
