package snippets

import (
	"phantom/dataLayer/dbModels"
	"phantom/dataLayer/uiModels/atoms"
	"strconv"
)

type ProductRailSnippet struct {
	Id        int64            `json:"id,omitempty"`
	Name      *atoms.TextData  `json:"name,omitempty"`
	ShortDesc *atoms.TextData  `json:"short_desc,omitempty"`
	Brand     *atoms.TextData  `json:"brand,omitempty"`
	Category  *atoms.TextData  `json:"category,omitempty"`
	Cost      *atoms.TextData  `json:"cost,omitempty"`
	Image     *atoms.ImageData `json:"image,omitempty"`
}

func MakeProductRailSnippet(
	product dbModels.Product,
	category dbModels.Category,
	brand dbModels.Brand,
) ProductRailSnippet {
	snippet := ProductRailSnippet{
		Id:        product.Id,
		Name:      &atoms.TextData{Text: product.Name},
		ShortDesc: &atoms.TextData{Text: product.ShortDescription},
		Brand:     &atoms.TextData{Text: brand.Name},
		Category:  &atoms.TextData{Text: category.Name},
		Cost:      &atoms.TextData{Text: strconv.FormatInt(product.Cost, 10)},
		Image:     &atoms.ImageData{Url: product.CardImage},
	}
	return snippet
}
