package snippets

import (
	"phantom/dataLayer/dbModels"
	"phantom/dataLayer/uiModels/atoms"
)

type CategoryRailSnippetData struct {
	Type  string          `json:"type,omitempty"`
	Id    int64           `json:"id,omitempty"`
	Name  *atoms.TextData `json:"name,omitempty"`
	Click interface{}     `json:"click,omitempty"`
}

func MakeCategoryRailSnippet(category dbModels.Category) CategoryRailSnippetData {
	snippet := CategoryRailSnippetData{
		Type: CategoryRailSnippet,
		Id:   category.Id,
		Name: &atoms.TextData{Text: category.Name},
		Click: atoms.CategoryClickData{
			Type:       atoms.ClickTypeOpenCategory,
			CategoryId: category.Id,
		},
	}
	return snippet
}
