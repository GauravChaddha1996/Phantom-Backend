package apiCommons

import (
	"phantom/dataLayer/uiModels/snippets"
)

func LogApiError(data ApiErrorLogData) {

}

func ToBaseSnippets(arr ...interface{}) *[]snippets.BaseSnippet {
	baseSnippets := make([]snippets.BaseSnippet, len(arr))
	for index, element := range arr {
		baseSnippets[index] = element.(snippets.BaseSnippet)
	}
	return &baseSnippets
}

func AppendIgnoringNils(
	slice []*snippets.SnippetSectionData, elems ...*snippets.SnippetSectionData,
) []*snippets.SnippetSectionData {
	for _, elem := range elems {
		if elem != nil {
			slice = append(slice, elem)
		}
	}
	return slice
}
