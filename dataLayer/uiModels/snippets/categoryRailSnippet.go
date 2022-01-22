package snippets

import (
	"phantom/dataLayer/dbModels"
	"phantom/dataLayer/uiModels/atoms"
	"unicode"
)

type CategoryRailSnippetData struct {
	Type           string           `json:"type,omitempty"`
	Id             int64            `json:"id,omitempty"`
	FirstCharacter *atoms.TextData  `json:"first_character,omitempty"`
	Name           *atoms.TextData  `json:"name,omitempty"`
	BgColor        *atoms.ColorData `json:"bg_color,omitempty"`
	Click          interface{}      `json:"click,omitempty"`
}

func MakeCategoryRailSnippet(category dbModels.Category) CategoryRailSnippetData {
	firstCharacter := rune(category.Name[0])
	firstCharacterUpper := string(unicode.ToUpper(firstCharacter))
	snippet := CategoryRailSnippetData{
		Type: CategoryRailSnippet,
		Id:   category.Id,
		FirstCharacter: &atoms.TextData{
			Text: firstCharacterUpper,
		},
		Name: &atoms.TextData{Text: category.Name},
		Click: atoms.CategoryClickData{
			Type:       atoms.ClickTypeOpenCategory,
			CategoryId: category.Id,
		},
	}
	return snippet
}
