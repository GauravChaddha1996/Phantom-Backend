package snippets

import (
	"github.com/spf13/cast"
	"phantom/dataLayer/dbModels"
	"phantom/dataLayer/uiModels/atoms"
)

type ProductDualSnippetData struct {
	Type       string           `json:"type,omitempty"`
	Id         int64            `json:"id,omitempty"`
	Name       *atoms.TextData  `json:"name,omitempty"`
	ShortDesc  *atoms.TextData  `json:"short_desc,omitempty"`
	Brand      *atoms.TextData  `json:"brand,omitempty"`
	Cost       *atoms.TextData  `json:"cost,omitempty"`
	NewTagText *atoms.TextData  `json:"new_tag_text,omitempty"`
	Image      *atoms.ImageData `json:"image,omitempty"`
	Click      interface{}      `json:"click,omitempty"`
}

func MakeProductDualSnippet(
	product dbModels.Product,
	brand dbModels.Brand,
	isNew bool,
) ProductDualSnippetData {
	var newTagText *atoms.TextData
	if isNew {
		newTagText = &atoms.TextData{Text: "New"}
	}
	snippet := ProductDualSnippetData{
		Type:       ProductDualSnippet,
		Id:         product.Id,
		Name:       &atoms.TextData{Text: product.Name},
		ShortDesc:  &atoms.TextData{Text: product.ShortDescription},
		Brand:      &atoms.TextData{Text: "By " + brand.Name},
		Cost:       &atoms.TextData{Text: "₹" + cast.ToString(product.Cost)},
		NewTagText: newTagText,
		Image:      &atoms.ImageData{Url: product.CardImage},
		Click: atoms.ProductClickData{
			Type:      atoms.ClickTypeOpenProduct,
			ProductId: product.Id,
		},
	}
	return snippet
}
