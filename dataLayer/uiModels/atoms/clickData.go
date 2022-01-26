package atoms

const (
	ClickTypeOpenProduct  = "OPEN_PRODUCT"
	ClickTypeOpenCategory = "OPEN_CATEGORY"
	ClickTypeAddProduct   = "ADD_PRODUCT"
)

type CategoryClickData struct {
	Type          string     `json:"type,omitempty"`
	CategoryId    int64      `json:"category_id,omitempty"`
	CategoryColor *ColorData `json:"category_color,omitempty"`
}

type ProductClickData struct {
	Type      string `json:"type,omitempty"`
	ProductId int64  `json:"product_id,omitempty"`
}

type AddProductClickData struct {
	Type             string `json:"type,omitempty"`
	ProductId        int64  `json:"id,omitempty"`
	Name             string `json:"name,omitempty"`
	ShortDescription string `json:"short_desc,omitempty"`
	BrandAndCategory string `json:"brand_and_category,omitempty"`
	Image            string `json:"image,omitempty"`
}
