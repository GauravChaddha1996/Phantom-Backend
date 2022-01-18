package snippets

import (
	"fmt"
	"phantom/dataLayer/dbModels"
	"phantom/dataLayer/uiModels/atoms"
	"strings"
)

type TextSnippetData struct {
	Type           string            `json:"type,omitempty"`
	TextSectionArr []TextSectionData `json:"text_section_arr,omitempty"`
}

type TextSectionData struct {
	Title     *atoms.TextData `json:"title,omitempty"`
	Subtitle  *atoms.TextData `json:"subtitle,omitempty"`
	Subtitle2 *atoms.TextData `json:"subtitle2,omitempty"`
}

func MakeTextSectionSnippetFromPropertyMapping(
	propertyMapping *map[dbModels.Property][]dbModels.PropertyValue,
) TextSnippetData {
	var textSectionArr = make([]TextSectionData, 0)
	for property, propertyValueArr := range *propertyMapping {
		// Find combined property value
		propertyValueStringArr := make([]string, len(propertyValueArr))
		for index, propertyValue := range propertyValueArr {
			propertyValueStringArr[index] = propertyValue.Name
		}
		combinedPropertyValue := strings.Join(propertyValueStringArr, ", ")

		section := TextSectionData{
			Title: &atoms.TextData{
				Text:  fmt.Sprintf("%s", strings.Title(property.Name)),
				Color: &atoms.ColorData{Name: atoms.COLOR_GREY_900},
				Font:  &atoms.FontData{Style: atoms.FONT_MEDIUM_500},
			},
			Subtitle: &atoms.TextData{
				Text:  strings.Title(combinedPropertyValue),
				Color: &atoms.ColorData{Name: atoms.COLOR_GREY_700},
				Font:  &atoms.FontData{Style: atoms.FONT_MEDIUM_400},
			},
		}
		textSectionArr = append(textSectionArr, section)
	}
	snippet := TextSnippetData{
		Type:           TextSnippet,
		TextSectionArr: textSectionArr,
	}
	return snippet
}
