package section

import (
	"phantom/apis/apiCommons"
	"phantom/dataLayer/dbModels"
	"phantom/dataLayer/uiModels/atoms"
	"phantom/dataLayer/uiModels/snippets"
)

func MakeHeaderSection(category *dbModels.Category) *snippets.SnippetSectionData {
	nameTextSection := snippets.TextSectionData{Title: &atoms.TextData{Text: category.Name}}
	descTextSection := snippets.TextSectionData{Title: &atoms.TextData{Text: category.Description}}
	textSectionArr := []snippets.TextSectionData{nameTextSection, descTextSection}
	textSnippet := snippets.TextSectionSnippet{TextSectionArr: textSectionArr}
	return &snippets.SnippetSectionData{
		Type:       snippets.PageHeaderSection,
		HeaderData: nil,
		Snippets:   apiCommons.ToBaseSnippets(textSnippet),
	}
}
