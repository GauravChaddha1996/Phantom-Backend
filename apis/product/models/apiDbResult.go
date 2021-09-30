package models

import "phantom/dataLayer/dbModels"

type ApiDbResult struct {
	Product         *dbModels.Product
	ProductImages   []dbModels.ProductImage
	Category        *dbModels.Category
	Brand           *dbModels.Brand
	PropertyMapping *map[dbModels.Property][]dbModels.PropertyValue
}
