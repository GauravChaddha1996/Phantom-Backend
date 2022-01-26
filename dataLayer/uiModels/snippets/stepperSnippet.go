package snippets

import (
	"fmt"
	"phantom/dataLayer/dbModels"
	"phantom/dataLayer/uiModels/atoms"
)

type StepperSnippetData struct {
	Type  string                     `json:"type,omitempty"`
	Title *atoms.TextData            `json:"title,omitempty"`
	Click *atoms.AddProductClickData `json:"click,omitempty"`
}

func MakeStepperSnippet(product dbModels.Product, brand dbModels.Brand, category dbModels.Category) StepperSnippetData {
	return StepperSnippetData{
		Type:  StepperSnippet,
		Title: &atoms.TextData{Text: fmt.Sprintf("Add for ₹%d", product.Cost)},
		Click: &atoms.AddProductClickData{
			Type:             atoms.ClickTypeAddProduct,
			ProductId:        product.Id,
			Name:             product.Name,
			ShortDescription: product.ShortDescription,
			BrandAndCategory: MakeBrandAndCategoryText(brand, category),
			Image:            product.CardImage,
		},
	}
}
