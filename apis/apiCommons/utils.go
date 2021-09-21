package apiCommons

import (
	"encoding/json"
	"log"
	"phantom/dataLayer/uiModels/snippets"
)

func LogApiError(data ApiErrorLogData) {
	marshaledApiErrorLogData, err := json.Marshal(data)
	if err != nil {
		return
	}
	log.Println(string(marshaledApiErrorLogData))
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
