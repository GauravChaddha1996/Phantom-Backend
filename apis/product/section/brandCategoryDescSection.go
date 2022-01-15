package section

import (
	"phantom/apis/apiCommons"
	"phantom/apis/product/models"
	"phantom/dataLayer/uiModels/atoms"
	"phantom/dataLayer/uiModels/snippets"
)

func MakeLongDescSection(apiDbResult *models.ApiDbResult) *snippets.SnippetSectionData {
	descSection := snippets.TextSectionData{
		Title: &atoms.TextData{Text: apiDbResult.Product.LongDescription},
	}
	textSectionSnippet := snippets.TextSectionSnippet{TextSectionArr: []snippets.TextSectionData{descSection}}
	return &snippets.SnippetSectionData{
		Type:       snippets.TextSnippet,
		HeaderData: nil,
		Snippets:   apiCommons.ToBaseSnippets(textSectionSnippet),
	}
}
