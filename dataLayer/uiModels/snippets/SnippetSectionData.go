package snippets

import "phantom/dataLayer/uiModels/atoms"

const ProductRailSection = "product_rail_section"
const ProductFullSection = "product_full_section"
const ProductDualSection = "product_dual_section"
const CategoryRailSection = "category_rail_section"

type SnippetSectionData struct {
	Type       string
	HeaderData *SnippetSectionHeaderData
	Snippets   []BaseSnippet
}

type SnippetSectionHeaderData struct {
	Title       *atoms.TextData
	Subtitle    *atoms.TextData
	RightButton *atoms.ButtonData
}
