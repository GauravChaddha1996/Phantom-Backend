package models

import (
	"phantom/dataLayer/uiModels/snippets"
)

type FilterApiResponse struct {
	Status            string                         `json:"status,omitempty"`
	Message           string                         `json:"message,omitempty"`
	Snippets          []*snippets.SnippetSectionData `json:"snippets,omitempty"`
	FilterSheetUiData FilterSheetUiData              `json:"filter_sheet_ui_data,omitempty"`
	SortSheetUiData   SortSheetUiData                `json:"sort_sheet_ui_data,omitempty"`
}
