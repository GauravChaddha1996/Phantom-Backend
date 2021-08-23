package atoms

type TextData struct {
	Text  string     `json:"text,omitempty"`
	Color *ColorData `json:"color,omitempty"`
	Font  *FontData  `json:"font,omitempty"`
}
