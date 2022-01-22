package section

import (
	"phantom/apis/product/models"
	"phantom/dataLayer/uiModels/atoms"
	"phantom/dataLayer/uiModels/snippets"
)

func MakeLongDescSection(apiDbResult *models.ApiDbResult) *snippets.SnippetSectionData {
	longDescSection := snippets.TextSnippetData{
		Type: snippets.TextSnippet,
		Subtitle: &atoms.TextData{
			Text:  apiDbResult.Product.LongDescription,
			Font:  &atoms.FontData{Style: atoms.FontBodyLarge},
		},
	}
	return &snippets.SnippetSectionData{
		HeaderData: &snippets.SnippetSectionHeaderData{
			Title: &atoms.TextData{
				Text: "About the product",
			},
		},
		Snippets: []snippets.TextSnippetData{longDescSection},
	}
}
