package models

import "phantom/dataLayer/dbModels"

type ApiDbResult struct {
	ProductsMap   map[int64]*dbModels.Product
	CategoriesMap map[int64]*dbModels.Category
	BrandsMap     map[int64]*dbModels.Brand
}

func EmptyHomeApiDbResult() ApiDbResult {
	return ApiDbResult{
		map[int64]*dbModels.Product{},
		map[int64]*dbModels.Category{},
		map[int64]*dbModels.Brand{},
	}
}
