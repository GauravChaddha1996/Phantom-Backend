package atoms

type ButtonType = string

const ButtonTypeText = "text"

type ButtonData struct {
	Text TextData
	Type ButtonType
}
