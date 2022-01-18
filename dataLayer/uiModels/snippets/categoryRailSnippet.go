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
	firstCharacterLower := string(unicode.ToLower(firstCharacter))
	snippet := CategoryRailSnippetData{
		Type: CategoryRailSnippet,
		Id:   category.Id,
		FirstCharacter: &atoms.TextData{
			Text:  firstCharacterUpper,
			Color: &atoms.ColorData{Name: getCategoryRailSnippetFirstCharColor(firstCharacterLower)},
		},
		Name:    &atoms.TextData{Text: category.Name},
		BgColor: &atoms.ColorData{Name: getCategoryRailSnippetBgColor(firstCharacterLower)},
		Click: atoms.CategoryClickData{
			Type:       atoms.ClickTypeOpenCategory,
			CategoryId: category.Id,
		},
	}
	return snippet
}

func getCategoryRailSnippetBgColor(firstCharacter string) string {
	var bgColor string
	switch firstCharacter {
	case "a":
		bgColor = atoms.ColorPrimary
		break
	case "b":
		bgColor = atoms.ColorOutline
		break
	case "c":
		bgColor = atoms.ColorSecondary
		break
	case "d":
		bgColor = atoms.ColorTertiary
		break
	case "e":
		bgColor = atoms.ColorPrimary
		break
	case "f":
		bgColor = atoms.ColorOutline
		break
	case "g":
		bgColor = atoms.ColorSecondary
		break
	case "h":
		bgColor = atoms.ColorTertiary
		break
	case "i":
		bgColor = atoms.ColorPrimary
		break
	case "j":
		bgColor = atoms.ColorOutline
		break
	case "k":
		bgColor = atoms.ColorSecondary
		break
	case "l":
		bgColor = atoms.ColorTertiary
		break
	case "m":
		bgColor = atoms.ColorPrimary
		break
	case "n":
		bgColor = atoms.ColorOutline
		break
	case "o":
		bgColor = atoms.ColorSecondary
		break
	case "p":
		bgColor = atoms.ColorTertiary
		break
	case "q":
		bgColor = atoms.ColorPrimary
		break
	case "r":
		bgColor = atoms.ColorOutline
		break
	case "s":
		bgColor = atoms.ColorSecondary
		break
	case "t":
		bgColor = atoms.ColorTertiary
		break
	case "u":
		bgColor = atoms.ColorTertiary
		break
	case "v":
		bgColor = atoms.ColorPrimary
		break
	case "w":
		bgColor = atoms.ColorOutline
		break
	case "x":
		bgColor = atoms.ColorSecondary
		break
	case "y":
		bgColor = atoms.ColorTertiary
		break
	case "z":
		bgColor = atoms.ColorTertiary
		break
	}
	return bgColor
}

func getCategoryRailSnippetFirstCharColor(firstCharacter string) string {
	var bgColor string
	switch firstCharacter {
	case "a":
		bgColor = atoms.ColorOnPrimary
		break
	case "b":
		bgColor = atoms.ColorOnPrimary
		break
	case "c":
		bgColor = atoms.ColorOnSecondary
		break
	case "d":
		bgColor = atoms.ColorOnTertiary
		break
	case "e":
		bgColor = atoms.ColorOnPrimary
		break
	case "f":
		bgColor = atoms.ColorOnPrimary
		break
	case "g":
		bgColor = atoms.ColorOnSecondary
		break
	case "h":
		bgColor = atoms.ColorOnTertiary
		break
	case "i":
		bgColor = atoms.ColorOnPrimary
		break
	case "j":
		bgColor = atoms.ColorOnPrimary
		break
	case "k":
		bgColor = atoms.ColorOnSecondary
		break
	case "l":
		bgColor = atoms.ColorOnTertiary
		break
	case "m":
		bgColor = atoms.ColorOnPrimary
		break
	case "n":
		bgColor = atoms.ColorOnPrimary
		break
	case "o":
		bgColor = atoms.ColorOnSecondary
		break
	case "p":
		bgColor = atoms.ColorOnTertiary
		break
	case "q":
		bgColor = atoms.ColorOnPrimary
		break
	case "r":
		bgColor = atoms.ColorOnPrimary
		break
	case "s":
		bgColor = atoms.ColorOnSecondary
		break
	case "t":
		bgColor = atoms.ColorOnTertiary
		break
	case "u":
		bgColor = atoms.ColorOnTertiary
		break
	case "v":
		bgColor = atoms.ColorOnPrimary
		break
	case "w":
		bgColor = atoms.ColorOnPrimary
		break
	case "x":
		bgColor = atoms.ColorOnSecondary
		break
	case "y":
		bgColor = atoms.ColorOnTertiary
		break
	case "z":
		bgColor = atoms.ColorOnTertiary
		break
	}
	return bgColor
}
