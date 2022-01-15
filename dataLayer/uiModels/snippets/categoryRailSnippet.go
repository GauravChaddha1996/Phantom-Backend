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
	backPageColor := &atoms.ColorData{Name: getCategoryScreenBackPageColor(firstCharacterLower)}
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
			Type:          atoms.ClickTypeOpenCategory,
			CategoryId:    category.Id,
			CategoryColor: backPageColor,
		},
	}
	return snippet
}

func getCategoryRailSnippetBgColor(firstCharacter string) string {
	var bgColor string
	switch firstCharacter {
	case "a":
		bgColor = atoms.COLOR_BLUE_100
		break
	case "b":
		bgColor = atoms.COLOR_RED_100
		break
	case "c":
		bgColor = atoms.COLOR_YELLOW_100
		break
	case "d":
		bgColor = atoms.COLOR_GREEN_100
		break
	case "e":
		bgColor = atoms.COLOR_BLUE_100
		break
	case "f":
		bgColor = atoms.COLOR_RED_100
		break
	case "g":
		bgColor = atoms.COLOR_YELLOW_100
		break
	case "h":
		bgColor = atoms.COLOR_GREEN_100
		break
	case "i":
		bgColor = atoms.COLOR_BLUE_100
		break
	case "j":
		bgColor = atoms.COLOR_RED_100
		break
	case "k":
		bgColor = atoms.COLOR_YELLOW_100
		break
	case "l":
		bgColor = atoms.COLOR_GREEN_100
		break
	case "m":
		bgColor = atoms.COLOR_BLUE_100
		break
	case "n":
		bgColor = atoms.COLOR_RED_100
		break
	case "o":
		bgColor = atoms.COLOR_YELLOW_100
		break
	case "p":
		bgColor = atoms.COLOR_GREEN_100
		break
	case "q":
		bgColor = atoms.COLOR_BLUE_100
		break
	case "r":
		bgColor = atoms.COLOR_RED_100
		break
	case "s":
		bgColor = atoms.COLOR_YELLOW_100
		break
	case "t":
		bgColor = atoms.COLOR_GREEN_100
		break
	case "u":
		bgColor = atoms.COLOR_GREEN_100
		break
	case "v":
		bgColor = atoms.COLOR_BLUE_100
		break
	case "w":
		bgColor = atoms.COLOR_RED_100
		break
	case "x":
		bgColor = atoms.COLOR_YELLOW_100
		break
	case "y":
		bgColor = atoms.COLOR_GREEN_100
		break
	case "z":
		bgColor = atoms.COLOR_GREEN_100
		break
	}
	return bgColor
}

func getCategoryRailSnippetFirstCharColor(firstCharacter string) string {
	var bgColor string
	switch firstCharacter {
	case "a":
		bgColor = atoms.COLOR_BLUE_400
		break
	case "b":
		bgColor = atoms.COLOR_RED_400
		break
	case "c":
		bgColor = atoms.COLOR_YELLOW_400
		break
	case "d":
		bgColor = atoms.COLOR_GREEN_400
		break
	case "e":
		bgColor = atoms.COLOR_BLUE_400
		break
	case "f":
		bgColor = atoms.COLOR_RED_400
		break
	case "g":
		bgColor = atoms.COLOR_YELLOW_400
		break
	case "h":
		bgColor = atoms.COLOR_GREEN_400
		break
	case "i":
		bgColor = atoms.COLOR_BLUE_400
		break
	case "j":
		bgColor = atoms.COLOR_RED_400
		break
	case "k":
		bgColor = atoms.COLOR_YELLOW_400
		break
	case "l":
		bgColor = atoms.COLOR_GREEN_400
		break
	case "m":
		bgColor = atoms.COLOR_BLUE_400
		break
	case "n":
		bgColor = atoms.COLOR_RED_400
		break
	case "o":
		bgColor = atoms.COLOR_YELLOW_400
		break
	case "p":
		bgColor = atoms.COLOR_GREEN_400
		break
	case "q":
		bgColor = atoms.COLOR_BLUE_400
		break
	case "r":
		bgColor = atoms.COLOR_RED_400
		break
	case "s":
		bgColor = atoms.COLOR_YELLOW_400
		break
	case "t":
		bgColor = atoms.COLOR_GREEN_400
		break
	case "u":
		bgColor = atoms.COLOR_GREEN_400
		break
	case "v":
		bgColor = atoms.COLOR_BLUE_400
		break
	case "w":
		bgColor = atoms.COLOR_RED_400
		break
	case "x":
		bgColor = atoms.COLOR_YELLOW_400
		break
	case "y":
		bgColor = atoms.COLOR_GREEN_400
		break
	case "z":
		bgColor = atoms.COLOR_GREEN_400
		break
	}
	return bgColor
}

func getCategoryScreenBackPageColor(firstCharacter string) string {
	var bgColor string
	switch firstCharacter {
	case "a":
		bgColor = atoms.COLOR_BLUE_300
		break
	case "b":
		bgColor = atoms.COLOR_RED_300
		break
	case "c":
		bgColor = atoms.COLOR_YELLOW_300
		break
	case "d":
		bgColor = atoms.COLOR_GREEN_300
		break
	case "e":
		bgColor = atoms.COLOR_BLUE_300
		break
	case "f":
		bgColor = atoms.COLOR_RED_300
		break
	case "g":
		bgColor = atoms.COLOR_YELLOW_300
		break
	case "h":
		bgColor = atoms.COLOR_GREEN_300
		break
	case "i":
		bgColor = atoms.COLOR_BLUE_300
		break
	case "j":
		bgColor = atoms.COLOR_RED_300
		break
	case "k":
		bgColor = atoms.COLOR_YELLOW_300
		break
	case "l":
		bgColor = atoms.COLOR_GREEN_300
		break
	case "m":
		bgColor = atoms.COLOR_BLUE_300
		break
	case "n":
		bgColor = atoms.COLOR_RED_300
		break
	case "o":
		bgColor = atoms.COLOR_YELLOW_300
		break
	case "p":
		bgColor = atoms.COLOR_GREEN_300
		break
	case "q":
		bgColor = atoms.COLOR_BLUE_300
		break
	case "r":
		bgColor = atoms.COLOR_RED_300
		break
	case "s":
		bgColor = atoms.COLOR_YELLOW_300
		break
	case "t":
		bgColor = atoms.COLOR_GREEN_300
		break
	case "u":
		bgColor = atoms.COLOR_GREEN_300
		break
	case "v":
		bgColor = atoms.COLOR_BLUE_300
		break
	case "w":
		bgColor = atoms.COLOR_RED_300
		break
	case "x":
		bgColor = atoms.COLOR_YELLOW_300
		break
	case "y":
		bgColor = atoms.COLOR_GREEN_300
		break
	case "z":
		bgColor = atoms.COLOR_GREEN_300
		break
	}
	return bgColor
}
