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
			Color: &atoms.ColorData{Name: atoms.COLOR_GREY_700},
			Font:  &atoms.FontData{Style: atoms.FONT_REGULAR_600},
		},
	}
	textSectionSnippet := snippets.TextSnippetData{Type: snippets.TextSnippet, TextSectionArr: []snippets.TextSectionData{longDescSection}}
	return &snippets.SnippetSectionData{
		HeaderData: &snippets.SnippetSectionHeaderData{
			Title: &atoms.TextData{
				Text: "About the product",
				Font: &atoms.FontData{Style: atoms.FONT_SEMIBOLD_500},
			},
		},
		Snippets: []snippets.TextSnippetData{textSectionSnippet},
	}
}
