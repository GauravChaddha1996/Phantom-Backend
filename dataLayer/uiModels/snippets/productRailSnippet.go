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
	Image            *atoms.ImageData `json:"image,omitempty"`
	Click            interface{}      `json:"click,omitempty"`
}

func MakeProductRailSnippet(
	product dbModels.Product,
	category dbModels.Category,
	brand dbModels.Brand,
) ProductRailSnippetData {
	brandAndCategoryText, brandAndCategoryMarkdownConfig := MakeBrandAndCategoryText(brand, category, atoms.FontBodyLarge)
	snippet := ProductRailSnippetData{
		Type:             ProductRailSnippet,
		Id:               product.Id,
		Name:             &atoms.TextData{Text: product.Name},
		ShortDesc:        &atoms.TextData{Text: product.ShortDescription},
		BrandAndCategory: &atoms.TextData{Text: brandAndCategoryText, MarkdownConfig: &brandAndCategoryMarkdownConfig},
		Cost:             &atoms.TextData{Text: "₹" + cast.ToString(product.Cost)},
		Image:            &atoms.ImageData{Url: product.CardImage},
		Click: atoms.ProductClickData{
			Type:      atoms.ClickTypeOpenProduct,
			ProductId: product.Id,
		},
	}
	return snippet
}

func MakeBrandAndCategoryText(
	brand dbModels.Brand, category dbModels.Category, fontStyle string,
) (string, atoms.MarkdownConfig) {
	brandPrefix := "By "
	categoryPrefix := " In "
	brandText := brandPrefix + brand.Name
	categoryText := categoryPrefix + category.Name
	brandAndCategoryText := brandText + categoryText
	brandAndCategoryMarkdownConfig := atoms.MarkdownConfig{
		Enabled: true,
		Spans: []interface{}{
			atoms.MarkdownFontSpan{
				Type:  atoms.MK_FONT_SPAN,
				Style: fontStyle,
				Start: len(brandPrefix),
				End:   len(brandText),
			},
			atoms.MarkdownFontSpan{
				Type:  atoms.MK_FONT_SPAN,
				Style: fontStyle,
				Start: len(brandText) + len(categoryPrefix),
				End:   len(brandAndCategoryText),
			},
		},
	}
	return brandAndCategoryText, brandAndCategoryMarkdownConfig
}
