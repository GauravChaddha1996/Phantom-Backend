package apiCommons

import (
	"phantom/dataLayer/dbModels"
)

type ProductIdMap struct {
	data map[int64]bool
}

func NewProductIdMap(products map[int64]*dbModels.Product) *ProductIdMap {
	if products == nil {
		return nil
	}
	dataMap := map[int64]bool{}
	for key := range products {
		dataMap[key] = false
	}
	productIdMap := ProductIdMap{data: dataMap}
	return &productIdMap
}

func (productIdMap ProductIdMap) PutAllInt64s(iArr *[]int64) {
	for _, productId := range *iArr {
		productIdMap.data[productId] = true
	}
}

func (productIdMap ProductIdMap) Put(productId int64) {
	productIdMap.data[productId] = true
}

func (productIdMap ProductIdMap) Contains(productId int64) bool {
	return productIdMap.data[productId] == true
}

func (productIdMap ProductIdMap) RemainingProducts() *[]int64 {
	var allFalseIds []int64
	for productId, value := range productIdMap.data {
		if !value {
			allFalseIds = append(allFalseIds, productId)
		}
	}
	return &allFalseIds
}
