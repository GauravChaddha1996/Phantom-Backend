package snippets

import (
	"phantom/dataLayer/dbModels"
	"phantom/dataLayer/uiModels/atoms"
	"strconv"
)

type ProductFullSnippet struct {
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
) ProductFullSnippet {
	snippet := ProductFullSnippet{
		Id:       product.Id,
		Name:     &atoms.TextData{Text: product.Name},
		LongDesc: &atoms.TextData{Text: product.LongDescription},
		Brand:    &atoms.TextData{Text: brand.Name},
		Category: &atoms.TextData{Text: category.Name},
		Cost:     &atoms.TextData{Text: strconv.FormatInt(product.Cost, 10)},
		Image:    &atoms.ImageData{Url: product.CardImage},
	}
	return snippet
}
