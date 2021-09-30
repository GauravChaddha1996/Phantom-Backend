package snippets

import "phantom/dataLayer/uiModels/atoms"

const ProductRailSection = "product_rail_section"
const ProductFullSection = "product_full_section"
const ProductDualSection = "product_dual_section"
const CategoryRailSection = "category_rail_section"
const PageHeaderSection = "page_header_section"
const ImagePagerSection = "image_pager_section"
const TextSection = "text_section"
const StepperSection = "stepper_section"

type SnippetSectionData struct {
	Type       string                    `json:"type,omitempty"`
	HeaderData *SnippetSectionHeaderData `json:"header_data,omitempty"`
	Snippets   *[]BaseSnippet            `json:"snippets,omitempty"`
}

type SnippetSectionHeaderData struct {
	Title       *atoms.TextData   `json:"title,omitempty"`
	Subtitle    *atoms.TextData   `json:"subtitle,omitempty"`
	RightButton *atoms.ButtonData `json:"right_button,omitempty"`
}
