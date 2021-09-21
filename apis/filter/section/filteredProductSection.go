package section

import (
	"fmt"
	"phantom/apis/apiCommons"
	"phantom/apis/filter/models"
	"phantom/dataLayer/uiModels/atoms"
	"phantom/dataLayer/uiModels/snippets"
)

const filtered_product_section_header_singular = "Showing 1 product"
const filtered_product_section_header_plural = "Showing %#v products"

func GetFilteredProductSnippetSection(apiDbResult *models.ApiDbResult) snippets.SnippetSectionData {
	productDualSnippets := getProductDualSnippetsFromDbResult(apiDbResult)
	sectionHeader := getProductSectionHeader(len(productDualSnippets))

	return snippets.SnippetSectionData{
		Type: snippets.ProductDualSection,
		HeaderData: &snippets.SnippetSectionHeaderData{
			Title:       &atoms.TextData{Text: sectionHeader},
			Subtitle:    nil,
			RightButton: nil,
		},
		Snippets: apiCommons.ToBaseSnippets(productDualSnippets),
	}
}

func getProductDualSnippetsFromDbResult(apiDbResult *models.ApiDbResult) []snippets.ProductDualSnippet {
	var productDualSnippets []snippets.ProductDualSnippet
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
		header = filtered_product_section_header_singular
	} else {
		header = fmt.Sprintf(filtered_product_section_header_plural, numberOfFilteredProductsFound)
	}
	return header
}
