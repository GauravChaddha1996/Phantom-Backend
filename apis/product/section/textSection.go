package section

import (
	"phantom/apis/product/models"
	"phantom/dataLayer/uiModels/atoms"
	"phantom/dataLayer/uiModels/snippets"
	"sort"
)

func MakePropertyMappingSection(apiDbResult *models.ApiDbResult) *snippets.SnippetSectionData {
	propertyMappingSection := snippets.MakeTextSnippetDataArrFromPropertyMapping(apiDbResult.PropertyMapping)
	sort.SliceStable(propertyMappingSection, func(one, two int) bool {
		textSnippetDataOne := propertyMappingSection[one]
		textSnippetDataTwo := propertyMappingSection[two]
		return textSnippetDataOne.Title.Text < textSnippetDataTwo.Title.Text
	})
	return &snippets.SnippetSectionData{
		HeaderData: &snippets.SnippetSectionHeaderData{
			Title: &atoms.TextData{
				Text: "Specifications",
			},
		},
		Snippets: propertyMappingSection,
	}
}
