package snippets

import "phantom/dataLayer/uiModels/atoms"

type BaseSnippet interface{}

type ProductRailSnippet struct {
	Id        int64            `json:"id,omitempty"`
	Name      *atoms.TextData  `json:"name,omitempty"`
	ShortDesc *atoms.TextData  `json:"short_desc,omitempty"`
	Brand     *atoms.TextData  `json:"brand,omitempty"`
	Category  *atoms.TextData  `json:"category,omitempty"`
	Cost      *atoms.TextData  `json:"cost,omitempty"`
	Image     *atoms.ImageData `json:"image,omitempty"`
}

type ProductFullSnippet struct {
	Id       int64            `json:"id,omitempty"`
	Name     *atoms.TextData  `json:"name,omitempty"`
	LongDesc *atoms.TextData  `json:"long_desc,omitempty"`
	Brand    *atoms.TextData  `json:"brand,omitempty"`
	Category *atoms.TextData  `json:"category,omitempty"`
	Cost     *atoms.TextData  `json:"cost,omitempty"`
	Image    *atoms.ImageData `json:"image,omitempty"`
}

type ProductDualSnippet struct {
	Id        int64            `json:"id,omitempty"`
	Name      *atoms.TextData  `json:"name,omitempty"`
	ShortDesc *atoms.TextData  `json:"short_desc,omitempty"`
	Brand     *atoms.TextData  `json:"brand,omitempty"`
	Cost      *atoms.TextData  `json:"cost,omitempty"`
	Image     *atoms.ImageData `json:"image,omitempty"`
}

type CategoryRailSnippet struct {
	Id   int64           `json:"id,omitempty"`
	Name *atoms.TextData `json:"name,omitempty"`
}
