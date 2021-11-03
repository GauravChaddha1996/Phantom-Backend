package snippets

import (
	"github.com/spf13/cast"
	"phantom/dataLayer/dbModels"
	"phantom/dataLayer/uiModels/atoms"
)

type ProductRailSnippetData struct {
	Type      string           `json:"type,omitempty"`
	Id        int64            `json:"id,omitempty"`
	Name      *atoms.TextData  `json:"name,omitempty"`
	ShortDesc *atoms.TextData  `json:"short_desc,omitempty"`
	Brand     *atoms.TextData  `json:"brand,omitempty"`
	Category  *atoms.TextData  `json:"category,omitempty"`
	Cost      *atoms.TextData  `json:"cost,omitempty"`
	Image     *atoms.ImageData `json:"image,omitempty"`
	Click     interface{}      `json:"click,omitempty"`
}

func MakeProductRailSnippet(
	product dbModels.Product,
	category dbModels.Category,
	brand dbModels.Brand,
) ProductRailSnippetData {
	snippet := ProductRailSnippetData{
		Type:      ProductRailSnippet,
		Id:        product.Id,
		Name:      &atoms.TextData{Text: product.Name},
		ShortDesc: &atoms.TextData{Text: product.ShortDescription},
		Brand:     &atoms.TextData{Text: brand.Name},
		Category:  &atoms.TextData{Text: category.Name},
		Cost:      &atoms.TextData{Text: cast.ToString(product.Cost)},
		Image:     &atoms.ImageData{Url: product.CardImage},
		Click: atoms.ProductClickData{
			Type:      atoms.ClickTypeOpenProduct,
			ProductId: product.Id,
		},
	}
	return snippet
}
