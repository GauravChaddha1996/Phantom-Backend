package models

type HomeApiResponse struct {
	Status             string        `json:"status,omitempty"`
	Message            string        `json:"message,omitempty"`
	SnippetSectionList interface{} `json:"snippet_section_list,omitempty"`
}
