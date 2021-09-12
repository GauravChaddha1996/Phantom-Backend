package snippets

import "phantom/dataLayer/uiModels/atoms"

type CategoryRailSnippet struct {
	Id   int64           `json:"id,omitempty"`
	Name *atoms.TextData `json:"name,omitempty"`
}
