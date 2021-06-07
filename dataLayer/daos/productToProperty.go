package daos

import (
	"database/sql"
	"phantom/dataLayer/dbModels"
)

type ProductToPropertyDao struct {
	DB *sql.DB
}

func (dao ProductToPropertyDao) CreateProductToPropertyMapping(productToProperty dbModels.ProductToProperty) (*int64, error) {
	query := "Insert into product_to_property (product_id, property_id, value_id) values (?, ?, ?)"
	lastInsertId, err := prepareAndExecuteInsertQuery(dao.DB, query, productToProperty.ProductId, productToProperty.PropertyId, productToProperty.ValueId)
	if err != nil {
		return nil, err
	}
	return lastInsertId, nil
}

func (dao ProductToPropertyDao) ReadAllProductToPropertyMapping(productId int64) (*[]dbModels.ProductToProperty, error) {
	var mappingArr []dbModels.ProductToProperty
	query := "Select * from product_to_property where product_id = ?"

	rows, queryErr := prepareAndExecuteSelectQuery(dao.DB, query, productId)
	if queryErr != nil {
		return nil, queryErr
	}
	defer closeRows(rows)

	for rows.Next() {
		var mapping dbModels.ProductToProperty
		rowErr := rows.Scan(&mapping.ProductId, &mapping.PropertyId, &mapping.ValueId)
		if rowErr != nil {
			continue
		}
		mappingArr = append(mappingArr, mapping)
	}
	return &mappingArr, queryErr
}
