package atoms

type TextData struct {
	Text           string          `json:"text,omitempty"`
	Color          *ColorData      `json:"color,omitempty"`
	Font           *FontData       `json:"font,omitempty"`
	MarkdownConfig *MarkdownConfig `json:"markdown_config,omitempty"`
}

type MarkdownConfig struct {
	Enabled bool        `json:"enabled,omitempty"`
	Spans   interface{} `json:"spans,omitempty"`
}

const (
	MK_FONT_SPAN = "font"
)

type MarkdownFontSpan struct {
	Type  string `json:"type,omitempty"`
	Style string `json:"style,omitempty"`
	Start int    `json:"start,omitempty"`
	End   int    `json:"end,omitempty"`
}
