package atoms

const (
	ClickTypeOpenProduct  = "OPEN_PRODUCT"
	ClickTypeOpenCategory = "OPEN_CATEGORY"
)

type CategoryClickData struct {
	Type       string `json:"type,omitempty"`
	CategoryId int64  `json:"category_id,omitempty"`
}

type ProductClickData struct {
	Type      string `json:"type,omitempty"`
	ProductId int64  `json:"product_id,omitempty"`
}
