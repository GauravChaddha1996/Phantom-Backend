package section

import (
	"fmt"
	"phantom/apis/filter/models"
	"phantom/dataLayer/uiModels/snippets"
)

const filteredProductSectionHeaderSingular = "Showing 1 product"
const filteredProductSectionHeaderPlural = "Showing %#v products"

func GetFilteredProductSnippetSection(apiDbResult *models.ApiDbResult) (snippets.SnippetSectionData, string) {
	productDualSnippets := getProductDualSnippetsFromDbResult(apiDbResult)
	sectionHeader := getProductSectionHeader(len(productDualSnippets))

	return snippets.SnippetSectionData{
		Type: snippets.ProductDualSnippet,
		Snippets: &productDualSnippets,
	}, sectionHeader
}

func getProductDualSnippetsFromDbResult(apiDbResult *models.ApiDbResult) []snippets.ProductDualSnippetData {
	var productDualSnippets []snippets.ProductDualSnippetData
	for _, product := range apiDbResult.ProductsList {
		brand := apiDbResult.BrandsMap[product.BrandId]
		snippet := snippets.MakeProductDualSnippet(product, brand)
		productDualSnippets = append(productDualSnippets, snippet)
	}
	return productDualSnippets
}

func getProductSectionHeader(numberOfFilteredProductsFound int) string {
	var header string
	if numberOfFilteredProductsFound == 1 {
		header = filteredProductSectionHeaderSingular
	} else {
		header = fmt.Sprintf(filteredProductSectionHeaderPlural, numberOfFilteredProductsFound)
	}
	return header
}
