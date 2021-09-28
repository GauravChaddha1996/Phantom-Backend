package models

import (
	"phantom/dataLayer/dbModels"
	"strings"
)

const (
	Sort_Newly_Added       = iota
	Sort_Price_Low_To_High = iota
	Sort_Price_High_To_Low = iota
	Sort_A_Z               = iota
)

type SortMethod struct {
	Id      int64
	Title   string
	Compare func(productOne dbModels.Product, productTwo dbModels.Product) bool
}

var AllSortMethods = []SortMethod{
	newlyAddedSortMethod,
	priceLowToHighSortMethod,
	priceHighToLowSortMethod,
	aZSortMethod,
}

func FindSortMethod(sortId int64) SortMethod {
	for _, method := range AllSortMethods {
		if method.Id == sortId {
			return method
		}
	}
	return newlyAddedSortMethod
}

var newlyAddedSortMethod = SortMethod{
	Id:    Sort_Newly_Added,
	Title: "Newly added",
	Compare: func(productOne dbModels.Product, productTwo dbModels.Product) bool {
		return productOne.CreatedAt.After(*productTwo.CreatedAt)
	},
}

var priceLowToHighSortMethod = SortMethod{
	Id:    Sort_Price_Low_To_High,
	Title: "Price low to high",
	Compare: func(productOne dbModels.Product, productTwo dbModels.Product) bool {
		return productOne.Cost < productTwo.Cost
	},
}

var priceHighToLowSortMethod = SortMethod{
	Id:    Sort_Price_High_To_Low,
	Title: "Price high to low",
	Compare: func(productOne dbModels.Product, productTwo dbModels.Product) bool {
		return productOne.Cost > productTwo.Cost
	},
}

var aZSortMethod = SortMethod{
	Id:    Sort_A_Z,
	Title: "Alphabetically",
	Compare: func(productOne dbModels.Product, productTwo dbModels.Product) bool {
		return strings.Compare(productOne.Name, productTwo.Name) < 0
	},
}
