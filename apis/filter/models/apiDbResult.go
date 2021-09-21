package models

import "phantom/dataLayer/dbModels"

type ApiDbResult struct {
	ProductsList               []dbModels.Product
	BrandsMap                  map[int64]dbModels.Brand
	PropertyToPropertyValueMap map[dbModels.Property][]dbModels.PropertyValue
}
