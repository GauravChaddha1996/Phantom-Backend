package models

import "phantom/dataLayer/uiModels/snippets"

type HomeApiResponse struct {
	Status   string                        `json:"status,omitempty"`
	Message  string                        `json:"message,omitempty"`
	Snippets []snippets.SnippetSectionData `json:"snippets,omitempty"`
}
