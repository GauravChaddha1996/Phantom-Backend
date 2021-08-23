package snippets

import "phantom/dataLayer/uiModels/atoms"

type BaseSnippet interface{}

type ProductRailSnippet struct {
	Id        string
	Name      *atoms.TextData
	ShortDesc *atoms.TextData
	Brand     *atoms.TextData
	Category  *atoms.TextData
	Cost      *atoms.TextData
	Image     *atoms.ImageData
}

type ProductFullSnippet struct {
	Id       string
	Name     *atoms.TextData
	LongDesc *atoms.TextData
	Brand    *atoms.TextData
	Category *atoms.TextData
	Cost     *atoms.TextData
	Image    *atoms.ImageData
}

type ProductDualSnippet struct {
	Id        string
	Name      *atoms.TextData
	ShortDesc *atoms.TextData
	Brand     *atoms.TextData
	Cost      *atoms.TextData
	Image     *atoms.ImageData
}

type CategoryRailSnippet struct {
	Id    string
	Name  *atoms.TextData
	Image *atoms.ImageData
}
