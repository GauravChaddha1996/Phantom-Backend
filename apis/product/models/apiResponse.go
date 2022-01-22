package models

import (
	"phantom/dataLayer/uiModels/snippets"
)

type ProductApiResponse struct {
	Status         string                         `json:"status,omitempty"`
	Message        string                         `json:"message,omitempty"`
	Snippets       []*snippets.SnippetSectionData `json:"snippets,omitempty"`
	StepperSnippet snippets.StepperSnippetData    `json:"stepper_snippet,omitempty"`
}
