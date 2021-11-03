package snippets

import (
	"phantom/dataLayer/dbModels"
	"phantom/dataLayer/uiModels/atoms"
)

type CategoryRailSnippetData struct {
	Type string          `json:"type,omitempty"`
	Id   int64           `json:"id,omitempty"`
	Name *atoms.TextData `json:"name,omitempty"`
}

func MakeCategoryRailSnippet(category dbModels.Category) CategoryRailSnippetData {
	snippet := CategoryRailSnippetData{
		Type: CategoryRailSnippet,
		Id:   category.Id,
		Name: &atoms.TextData{Text: category.Name},
	}
	return snippet
}
