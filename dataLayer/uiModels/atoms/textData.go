package atoms

type TextData struct {
	Text           string          `json:"text,omitempty"`
	Color          *ColorData      `json:"color,omitempty"`
	Font           *FontData       `json:"font,omitempty"`
	MarkdownConfig *MarkdownConfig `json:"markdown_config,omitempty"`
}

type MarkdownConfig struct {
	Enabled bool `json:"enabled,omitempty"`
}
