package snippets

import "phantom/dataLayer/uiModels/atoms"

const ProductRailSnippet = "ProductRailSnippet"
const ProductFullSnippet = "ProductFullSnippet"
const CategoryRailSnippet = "CategoryRailSnippet"
const ProductDualSnippet = "ProductDualSnippet"
const TextSnippet = "TextSnippet"
const ImagePagerSection = "image_pager_section"
const StepperSection = "stepper_section"

type SnippetSectionData struct {
	Type       string                    `json:"type,omitempty"`
	HeaderData *SnippetSectionHeaderData `json:"header_data,omitempty"`
	Snippets   interface{}               `json:"snippets,omitempty"`
}

type SnippetSectionHeaderData struct {
	Title       *atoms.TextData   `json:"title,omitempty"`
	Subtitle    *atoms.TextData   `json:"subtitle,omitempty"`
	RightButton *atoms.ButtonData `json:"right_button,omitempty"`
}
