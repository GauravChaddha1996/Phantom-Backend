package section

import (
	"phantom/apis/product/models"
	"phantom/dataLayer/uiModels/atoms"
	"phantom/dataLayer/uiModels/snippets"
)

func MakePropertyMappingSection(apiDbResult *models.ApiDbResult) *snippets.SnippetSectionData {
	propertyMappingSection := snippets.MakeTextSnippetDataArrFromPropertyMapping(apiDbResult.PropertyMapping)
	return &snippets.SnippetSectionData{
		HeaderData: &snippets.SnippetSectionHeaderData{
			Title: &atoms.TextData{
				Text: "Specifications",
			},
		},
		Snippets: propertyMappingSection,
	}
}
