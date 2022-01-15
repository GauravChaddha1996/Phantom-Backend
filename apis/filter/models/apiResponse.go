package models

import (
	"phantom/dataLayer/uiModels/atoms"
	"phantom/dataLayer/uiModels/snippets"
)

type FilterApiResponse struct {
	Status               string                         `json:"status,omitempty"`
	Message              string                         `json:"message,omitempty"`
	PageTitle            atoms.TextData                 `json:"page_title,omitempty"`
	SnippetSectionHeader atoms.TextData                 `json:"snippet_section_header,omitempty"`
	SnippetSectionList   []*snippets.SnippetSectionData `json:"snippet_section_list,omitempty"`
	FilterSheetUiData    FilterSheetUiData              `json:"filter_sheet_ui_data,omitempty"`
	SortSheetUiData      SortSheetUiData                `json:"sort_sheet_ui_data,omitempty"`
}
