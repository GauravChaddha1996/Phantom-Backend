package section

import (
	"phantom/apis/product/models"
	"phantom/dataLayer/uiModels/snippets"
)

func MakeImagesSection(apiDbResult *models.ApiDbResult) *snippets.SnippetSectionData {
	imageSnippets := snippets.MakeProductImagesPagerSnippet(apiDbResult.ProductImages)
	return &snippets.SnippetSectionData{
		HeaderData: nil,
		Snippets:   []snippets.ImagePagerSnippetData{imageSnippets},
	}
}
