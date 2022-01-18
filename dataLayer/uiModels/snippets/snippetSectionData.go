package snippets

import "phantom/dataLayer/uiModels/atoms"

const ProductRailSnippet = "ProductRailSnippet"
const ProductFullSnippet = "ProductFullSnippet"
const CategoryRailSnippet = "CategoryRailSnippet"
const ProductDualSnippet = "ProductDualSnippet"
const TextSnippet = "TextSnippet"
const ImagePagerSnippet = "ImagePagerSnippet"
const StepperSnippet = "StepperSnippet"

type SnippetSectionData struct {
	HeaderData *SnippetSectionHeaderData `json:"header_data,omitempty"`
	Snippets   interface{}               `json:"snippets,omitempty"`
}

type SnippetSectionHeaderData struct {
	Title       *atoms.TextData   `json:"title,omitempty"`
	Subtitle    *atoms.TextData   `json:"subtitle,omitempty"`
	RightButton *atoms.ButtonData `json:"right_button,omitempty"`
}
