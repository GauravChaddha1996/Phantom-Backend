package snippets

import (
	"github.com/spf13/cast"
	"phantom/dataLayer/dbModels"
	"phantom/dataLayer/uiModels/atoms"
)

type ProductRailSnippetData struct {
	Type             string           `json:"type,omitempty"`
	Id               int64            `json:"id,omitempty"`
	Name             *atoms.TextData  `json:"name,omitempty"`
	ShortDesc        *atoms.TextData  `json:"short_desc,omitempty"`
	BrandAndCategory *atoms.TextData  `json:"brand_and_category,omitempty"`
	Cost             *atoms.TextData  `json:"cost,omitempty"`
	NewTagText       *atoms.TextData  `json:"new_tag_text,omitempty"`
	Image            *atoms.ImageData `json:"image,omitempty"`
	Click            interface{}      `json:"click,omitempty"`
}

func MakeProductRailSnippet(
	product dbModels.Product,
	category dbModels.Category,
	brand dbModels.Brand,
	isNew bool,
) ProductRailSnippetData {
	brandAndCategoryText := MakeBrandAndCategoryText(brand, category)
	var newTagText *atoms.TextData
	if isNew {
		newTagText = &atoms.TextData{Text: "New"}
	}
	snippet := ProductRailSnippetData{
		Type:             ProductRailSnippet,
		Id:               product.Id,
		Name:             &atoms.TextData{Text: product.Name},
		ShortDesc:        &atoms.TextData{Text: product.ShortDescription},
		BrandAndCategory: &atoms.TextData{Text: brandAndCategoryText},
		Cost:             &atoms.TextData{Text: "₹" + cast.ToString(product.Cost)},
		NewTagText:       newTagText,
		Image:            &atoms.ImageData{Url: product.CardImage},
		Click: atoms.ProductClickData{
			Type:      atoms.ClickTypeOpenProduct,
			ProductId: product.Id,
		},
	}
	return snippet
}

func MakeBrandAndCategoryText(brand dbModels.Brand, category dbModels.Category) string {
	brandPrefix := "By "
	categoryPrefix := " In "
	brandText := brandPrefix + brand.Name
	categoryText := categoryPrefix + category.Name
	brandAndCategoryText := brandText + categoryText
	return brandAndCategoryText
}
