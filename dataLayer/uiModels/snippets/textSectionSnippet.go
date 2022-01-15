package snippets

import (
	"fmt"
	"phantom/dataLayer/dbModels"
	"phantom/dataLayer/uiModels/atoms"
	"strings"
)

type TextSectionSnippet struct {
	Type           string            `json:"type,omitempty"`
	TextSectionArr []TextSectionData `json:"text_section_arr,omitempty"`
}

type TextSectionData struct {
	Title    *atoms.TextData `json:"title,omitempty"`
	Subtitle *atoms.TextData `json:"subtitle,omitempty"`
}

func MakeTextSectionSnippetFromPropertyMapping(
	propertyMapping *map[dbModels.Property][]dbModels.PropertyValue,
) TextSectionSnippet {
	var textSectionArr = make([]TextSectionData, 0)
	for property, propertyValueArr := range *propertyMapping {
		// Find combined property value
		propertyValueStringArr := make([]string, len(propertyValueArr))
		for index, propertyValue := range propertyValueArr {
			propertyValueStringArr[index] = propertyValue.Name
		}
		combinedPropertyValue := strings.Join(propertyValueStringArr, ", ")

		section := TextSectionData{
			&atoms.TextData{
				Text: fmt.Sprintf("%s:", property.Name),
			},
			&atoms.TextData{
				Text: combinedPropertyValue,
			},
		}
		textSectionArr = append(textSectionArr, section)
	}
	snippet := TextSectionSnippet{
		Type:           TextSnippet,
		TextSectionArr: textSectionArr,
	}
	return snippet
}
