package models

import (
	"phantom/apis/apiCommons"
	"phantom/dataLayer/dbModels"
)

type ApiDbResult struct {
	ProductsMap      map[int64]*dbModels.Product
	CategoriesMap    map[int64]*dbModels.Category
	BrandsMap        map[int64]*dbModels.Brand
	NewProductIdsMap *apiCommons.ProductIdMap
}

func EmptyHomeApiDbResult() ApiDbResult {
	return ApiDbResult{
		map[int64]*dbModels.Product{},
		map[int64]*dbModels.Category{},
		map[int64]*dbModels.Brand{},
		apiCommons.NewProductIdMap(map[int64]*dbModels.Product{}),
	}
}
