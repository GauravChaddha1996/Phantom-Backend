package snippets

import (
	"phantom/dataLayer/dbModels"
	"phantom/dataLayer/uiModels/atoms"
)

type ImagePagerSnippetData struct {
	Type   string            `json:"type,omitempty"`
	Images []atoms.ImageData `json:"images,omitempty"`
}

func MakeProductImagesPagerSnippet(
	productImages []dbModels.ProductImage,
) ImagePagerSnippetData {
	var images = make([]atoms.ImageData, len(productImages))
	for i, productImage := range productImages {
		images[i] = atoms.ImageData{Url: productImage.Url}
	}
	snippet := ImagePagerSnippetData{
		Type:   ImagePagerSnippet,
		Images: images,
	}
	return snippet
}
