package snippets

import (
	"github.com/spf13/cast"
	"phantom/dataLayer/dbModels"
	"phantom/dataLayer/uiModels/atoms"
)

type ProductDualSnippet struct {
	Id        int64            `json:"id,omitempty"`
	Name      *atoms.TextData  `json:"name,omitempty"`
	ShortDesc *atoms.TextData  `json:"short_desc,omitempty"`
	Brand     *atoms.TextData  `json:"brand,omitempty"`
	Cost      *atoms.TextData  `json:"cost,omitempty"`
	Image     *atoms.ImageData `json:"image,omitempty"`
}

func MakeProductDualSnippet(
	product dbModels.Product,
	brand dbModels.Brand,
) ProductDualSnippet {
	snippet := ProductDualSnippet{
		Id:        product.Id,
		Name:      &atoms.TextData{Text: product.Name},
		ShortDesc: &atoms.TextData{Text: product.ShortDescription},
		Brand:     &atoms.TextData{Text: brand.Name},
		Cost:      &atoms.TextData{Text: cast.ToString(product.Cost)},
		Image:     &atoms.ImageData{Url: product.CardImage},
	}
	return snippet
}
