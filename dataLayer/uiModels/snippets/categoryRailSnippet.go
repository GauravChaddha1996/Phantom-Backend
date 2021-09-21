package snippets

import (
	"phantom/dataLayer/dbModels"
	"phantom/dataLayer/uiModels/atoms"
)

type CategoryRailSnippet struct {
	Id   int64           `json:"id,omitempty"`
	Name *atoms.TextData `json:"name,omitempty"`
}

func MakeCategoryRailSnippet(category dbModels.Category) CategoryRailSnippet {
	snippet := CategoryRailSnippet{
		Id:   category.Id,
		Name: &atoms.TextData{Text: category.Name},
	}
	return snippet
}
