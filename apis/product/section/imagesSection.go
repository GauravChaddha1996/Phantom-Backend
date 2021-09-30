package section

import (
	"phantom/apis/apiCommons"
	"phantom/apis/product/models"
	"phantom/dataLayer/uiModels/snippets"
)

func MakeImagesSection(apiDbResult *models.ApiDbResult) *snippets.SnippetSectionData {
	imageSnippets := snippets.MakeProductImagesPagerSnippet(apiDbResult.ProductImages)
	return &snippets.SnippetSectionData{
		Type:       snippets.ImagePagerSection,
		HeaderData: nil,
		Snippets:   apiCommons.ToBaseSnippets(imageSnippets),
	}
}
