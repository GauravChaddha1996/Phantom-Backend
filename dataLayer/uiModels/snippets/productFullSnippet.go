package snippets

import (
	"github.com/spf13/cast"
	"phantom/dataLayer/dbModels"
	"phantom/dataLayer/uiModels/atoms"
)

type ProductFullSnippetData struct {
	Type             string           `json:"type,omitempty"`
	Id               int64            `json:"id,omitempty"`
	Name             *atoms.TextData  `json:"name,omitempty"`
	LongDesc         *atoms.TextData  `json:"long_desc,omitempty"`
	BrandAndCategory *atoms.TextData  `json:"brand_and_category,omitempty"`
	Cost             *atoms.TextData  `json:"cost,omitempty"`
	Image            *atoms.ImageData `json:"image,omitempty"`
	Click            interface{}      `json:"click,omitempty"`
}

func MakeProductFullSnippet(
	product dbModels.Product,
	category dbModels.Category,
	brand dbModels.Brand,
) ProductFullSnippetData {
	brandAndCategoryText, brandAndCategoryMarkdownConfig := MakeBrandAndCategoryText(brand, category)
	snippet := ProductFullSnippetData{
		Type:             ProductFullSnippet,
		Id:               product.Id,
		Name:             &atoms.TextData{Text: product.Name},
		LongDesc:         &atoms.TextData{Text: product.LongDescription + "Soft cotton shirt made by well paid hard Soft cotton shirt made by well paid hard Soft cotton shirt made by well paid hard Soft cotton shirt made by well paid hard "},
		BrandAndCategory: &atoms.TextData{Text: brandAndCategoryText, MarkdownConfig: &brandAndCategoryMarkdownConfig},
		Cost:             &atoms.TextData{Text: cast.ToString(product.Cost)},
		Image:            &atoms.ImageData{Url: product.CardImage},
		Click: atoms.ProductClickData{
			Type:      atoms.ClickTypeOpenProduct,
			ProductId: product.Id,
		},
	}
	return snippet
}
