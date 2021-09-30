package section

import (
	"phantom/apis/apiCommons"
	"phantom/apis/product/models"
	"phantom/dataLayer/uiModels/snippets"
)

func MakePropertyMappingSection(apiDbResult *models.ApiDbResult) *snippets.SnippetSectionData {
	propertyMappingSection := snippets.MakeTextSectionSnippetFromPropertyMapping(apiDbResult.PropertyMapping)
	return &snippets.SnippetSectionData{
		Type:       snippets.TextSection,
		HeaderData: nil,
		Snippets:   apiCommons.ToBaseSnippets(propertyMappingSection),
	}
}
