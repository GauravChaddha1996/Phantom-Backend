package section

import (
	"phantom/apis/product/models"
	"phantom/dataLayer/uiModels/atoms"
	"phantom/dataLayer/uiModels/snippets"
)

func MakeHeaderSection(apiDbResult *models.ApiDbResult) *snippets.SnippetSectionData {
	product := apiDbResult.Product
	brandCategoryText, markdownConfig := snippets.MakeBrandAndCategoryText(*apiDbResult.Brand, *apiDbResult.Category, atoms.FONT_SEMIBOLD_600)

	nameAndDescTextSection := snippets.TextSectionData{
		Title: &atoms.TextData{
			Text:  product.Name,
			Color: &atoms.ColorData{Name: atoms.COLOR_GREY_900},
			Font:  &atoms.FontData{Style: atoms.FONT_SEMIBOLD_600},
		},
		Subtitle: &atoms.TextData{
			Text:  product.ShortDescription,
			Color: &atoms.ColorData{Name: atoms.COLOR_GREY_500},
			Font:  &atoms.FontData{Style: atoms.FONT_REGULAR_400},
		},
		Subtitle2: &atoms.TextData{
			Text:           brandCategoryText,
			Color:          &atoms.ColorData{Name: atoms.COLOR_GREY_800},
			Font:           &atoms.FontData{Style: atoms.FONT_MEDIUM_600},
			MarkdownConfig: &markdownConfig,
		},
	}
	textSectionArr := []snippets.TextSectionData{nameAndDescTextSection}
	textSnippet := snippets.TextSnippetData{Type: snippets.TextSnippet, TextSectionArr: textSectionArr}
	return &snippets.SnippetSectionData{
		Snippets: []snippets.TextSnippetData{textSnippet},
	}
}
