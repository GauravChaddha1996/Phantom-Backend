package snippets

import (
	"github.com/spf13/cast"
	"phantom/dataLayer/dbModels"
	"phantom/dataLayer/uiModels/atoms"
)

type ProductFullSnippetData struct {
	Type     string           `json:"type,omitempty"`
	Id       int64            `json:"id,omitempty"`
	Name     *atoms.TextData  `json:"name,omitempty"`
	LongDesc *atoms.TextData  `json:"long_desc,omitempty"`
	Brand    *atoms.TextData  `json:"brand,omitempty"`
	Category *atoms.TextData  `json:"category,omitempty"`
	Cost     *atoms.TextData  `json:"cost,omitempty"`
	Image    *atoms.ImageData `json:"image,omitempty"`
}

func MakeProductFullSnippet(
	product dbModels.Product,
	category dbModels.Category,
	brand dbModels.Brand,
) ProductFullSnippetData {
	snippet := ProductFullSnippetData{
		Type:     ProductFullSnippet,
		Id:       product.Id,
		Name:     &atoms.TextData{Text: product.Name},
		LongDesc: &atoms.TextData{Text: product.LongDescription},
		Brand:    &atoms.TextData{Text: brand.Name},
		Category: &atoms.TextData{Text: category.Name},
		Cost:     &atoms.TextData{Text: cast.ToString(product.Cost)},
		Image:    &atoms.ImageData{Url: product.CardImage},
	}
	return snippet
}
