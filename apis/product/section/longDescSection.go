package section

import (
	"phantom/apis/product/models"
	"phantom/dataLayer/uiModels/atoms"
	"phantom/dataLayer/uiModels/snippets"
)

func MakeLongDescSection(apiDbResult *models.ApiDbResult) *snippets.SnippetSectionData {
	longDescSection := snippets.TextSectionData{
		Title: &atoms.TextData{
			Text:  apiDbResult.Product.LongDescription,
			Color: &atoms.ColorData{Name: atoms.ColorOnBackground},
			Font:  &atoms.FontData{Style: atoms.FontBodyMedium},
		},
	}
	textSectionSnippet := snippets.TextSnippetData{Type: snippets.TextSnippet, TextSectionArr: []snippets.TextSectionData{longDescSection}}
	return &snippets.SnippetSectionData{
		HeaderData: &snippets.SnippetSectionHeaderData{
			Title: &atoms.TextData{
				Text: "About the product",
			},
		},
		Snippets: []snippets.TextSnippetData{textSectionSnippet},
	}
}
