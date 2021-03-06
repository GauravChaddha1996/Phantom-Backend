package snippets

import (
	"fmt"
	"phantom/dataLayer/dbModels"
	"phantom/dataLayer/uiModels/atoms"
	"strings"
)
type TextSnippetData struct {
	Type            string          `json:"type,omitempty"`
	Title           *atoms.TextData `json:"title,omitempty"`
	Subtitle        *atoms.TextData `json:"subtitle,omitempty"`
	Subtitle2       *atoms.TextData `json:"subtitle2,omitempty"`
	BottomSeparator bool            `json:"bottom_separator,omitempty"`
}

func MakeTextSnippetDataArrFromPropertyMapping(
	propertyMapping *map[dbModels.Property][]dbModels.PropertyValue,
) []TextSnippetData {
	var textSectionArr = make([]TextSnippetData, 0)
	for property, propertyValueArr := range *propertyMapping {
		textSectionArr = append(textSectionArr, MakeTextSnippetDataFromPropertyValue(property, propertyValueArr))
	}
	return textSectionArr
}

func MakeTextSnippetDataFromPropertyValue(
	property dbModels.Property,
	propertyValueArr []dbModels.PropertyValue,
) TextSnippetData {
	// Find combined property value
	propertyValueStringArr := make([]string, len(propertyValueArr))
	for index, propertyValue := range propertyValueArr {
		propertyValueStringArr[index] = propertyValue.Name
	}
	combinedPropertyValue := strings.Join(propertyValueStringArr, ", ")

	return TextSnippetData{
		Type: TextSnippet,
		Title: &atoms.TextData{
			Text: fmt.Sprintf("%s", strings.Title(property.Name)),
			Font: &atoms.FontData{Style: atoms.FontTitleMedium},
		},
		Subtitle: &atoms.TextData{
			Text: strings.Title(combinedPropertyValue),
			Font: &atoms.FontData{Style: atoms.FontBodyLarge},
		},
	}
}
