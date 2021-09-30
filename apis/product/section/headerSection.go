package section

import (
	"fmt"
	"phantom/apis/apiCommons"
	"phantom/apis/product/models"
	"phantom/dataLayer/uiModels/atoms"
	"phantom/dataLayer/uiModels/snippets"
)

func MakeHeaderSection(apiDbResult *models.ApiDbResult) *snippets.SnippetSectionData {
	product := apiDbResult.Product
	brandCategoryText := fmt.Sprintf("By **%s** in **%s**", apiDbResult.Brand.Name, apiDbResult.Category.Name)

	nameTextSection := snippets.TextSectionData{Title: &atoms.TextData{Text: product.Name}}
	descTextSection := snippets.TextSectionData{Title: &atoms.TextData{Text: product.ShortDescription}}
	brandCategoryTextSection := snippets.TextSectionData{
		Title: &atoms.TextData{
			Text:           brandCategoryText,
			MarkdownConfig: &atoms.MarkdownConfig{Enabled: true},
		},
	}
	textSectionArr := []snippets.TextSectionData{nameTextSection, descTextSection, brandCategoryTextSection}
	textSnippet := snippets.TextSectionSnippet{TextSectionArr: textSectionArr}
	return &snippets.SnippetSectionData{
		Type:       snippets.PageHeaderSection,
		HeaderData: nil,
		Snippets:   apiCommons.ToBaseSnippets(textSnippet),
	}
}
