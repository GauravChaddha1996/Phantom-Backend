package snippets

import (
	"phantom/dataLayer/dbModels"
	"phantom/dataLayer/uiModels/atoms"
)

type ImagePagerSnippet struct {
	Images []atoms.ImageData `json:"images,omitempty"`
}

func MakeProductImagesPagerSnippet(
	productImages []dbModels.ProductImage,
) ImagePagerSnippet {
	var images = make([]atoms.ImageData, len(productImages))
	for i, productImage := range productImages {
		images[i] = atoms.ImageData{Url: productImage.Url}
	}
	snippet := ImagePagerSnippet{
		Images: images,
	}
	return snippet
}
