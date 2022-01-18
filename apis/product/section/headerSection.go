package section

import (
	"phantom/apis/product/models"
	"phantom/dataLayer/uiModels/atoms"
	"phantom/dataLayer/uiModels/snippets"
)

func MakeHeaderSection(apiDbResult *models.ApiDbResult) *snippets.SnippetSectionData {
	product := apiDbResult.Product
	brandCategoryText, markdownConfig := snippets.MakeBrandAndCategoryText(*apiDbResult.Brand, *apiDbResult.Category, atoms.FontBodyLarge)

	nameAndDescTextSection := snippets.TextSectionData{
		Title: &atoms.TextData{
			Text: product.Name,
		},
		Subtitle: &atoms.TextData{
			Text: product.ShortDescription,
		},
		Subtitle2: &atoms.TextData{
			Text:           brandCategoryText,
			MarkdownConfig: &markdownConfig,
		},
	}
	textSectionArr := []snippets.TextSectionData{nameAndDescTextSection}
	textSnippet := snippets.TextSnippetData{Type: snippets.TextSnippet, TextSectionArr: textSectionArr}
	return &snippets.SnippetSectionData{
		Snippets: []snippets.TextSnippetData{textSnippet},
	}
}
