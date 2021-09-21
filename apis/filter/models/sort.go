package models

const (
	Sort_Newly_Added       = iota
	Sort_Price_Low_To_High = iota
	Sort_Price_High_To_Low = iota
	Sort_Brand_Wise        = iota
	Sort_A_Z               = iota
)

type SortMethod struct {
	Id    int64
	Title string
}

var AllSortMethods = []SortMethod{
	{
		Id:    Sort_Newly_Added,
		Title: "Newly added",
	},
	{
		Id:    Sort_Price_Low_To_High,
		Title: "Price low to high",
	},
	{
		Id:    Sort_Price_High_To_Low,
		Title: "Price high to low",
	},
	{
		Id:    Sort_Brand_Wise,
		Title: "Brand-wise",
	},
	{
		Id:    Sort_A_Z,
		Title: "Alphabetically",
	},
}
