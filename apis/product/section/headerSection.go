package section

import (
	"phantom/apis/product/models"
	"phantom/dataLayer/uiModels/atoms"
	"phantom/dataLayer/uiModels/snippets"
)

func MakeHeaderSection(apiDbResult *models.ApiDbResult) *snippets.SnippetSectionData {
	product := apiDbResult.Product
	brandCategoryText := snippets.MakeBrandAndCategoryText(*apiDbResult.Brand, *apiDbResult.Category)

	nameAndDescTextSection := snippets.TextSnippetData{
		Type: snippets.TextSnippet,
		Title: &atoms.TextData{
			Text: product.Name,
		},
		Subtitle: &atoms.TextData{
			Text: product.ShortDescription,
		},
		Subtitle2: &atoms.TextData{
			Text: brandCategoryText,
		},
		BottomSeparator: true,
	}
	return &snippets.SnippetSectionData{
		Snippets: []snippets.TextSnippetData{nameAndDescTextSection},
	}
}
