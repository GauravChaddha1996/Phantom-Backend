package atoms

const (
	ClickTypeOpenCategory = "open_category"
)

type ClickData struct {
	Type string      `json:"type,omitempty"`
	Data interface{} `json:"data,omitempty"`
}

type CategoryClickData struct {
	CategoryId int64 `json:"category_id,omitempty"`
}
