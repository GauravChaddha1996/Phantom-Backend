package apiCommons

import (
	"phantom/dataLayer/uiModels/snippets"
)

func ToBaseSnippets(arr ...interface{}) *[]snippets.BaseSnippet {
	baseSnippets := make([]snippets.BaseSnippet, len(arr))
	for index, element := range arr {
		baseSnippets[index] = element.(snippets.BaseSnippet)
	}
	return &baseSnippets
}
