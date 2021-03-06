package atoms

type ButtonType = string

const ButtonTypeText = "text"

type ButtonData struct {
	Text  TextData    `json:"text,omitempty"`
	Type  ButtonType  `json:"type,omitempty"`
	Click interface{} `json:"click,omitempty"`
}
