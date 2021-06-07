package databasDaos

import (
	"database/sql"
	"phantom/dataLayer/dbModels"
)

type ProductDao struct {
	DB *sql.DB
}

func (dao ProductDao) CreateProduct(product dbModels.Product) (*int64, error) {
	query := "Insert into product " +
		"(id, brand_id, category_id, name, long_description, short_description, cost, card_image) " +
		"values (?, ?, ?, ?, ?, ?, ?, ?)"
	lastInsertId, err := prepareAndExecuteInsertQuery(dao.DB, query, product.Id, product.BrandId,
		product.CategoryId, product.Name, product.LongDescription,
		product.ShortDescription, product.Cost, product.CardImage)
	if err != nil {
		return nil, err
	}
	return lastInsertId, nil
}

func (dao ProductDao) ReadProduct(id int64) (*dbModels.Product, error) {
	var product dbModels.Product
	query := "Select * from product where id = ? limit 1"

	row, err := prepareAndExecuteSingleRowSelectQuery(dao.DB, query, id)
	if err != nil {
		return nil, err
	}

	err = row.Scan(&product.Id, &product.BrandId, &product.CategoryId,
		&product.Name, &product.LongDescription,
		&product.ShortDescription, &product.Cost, &product.CardImage)
	if err != nil {
		return nil, err
	}
	return &product, err
}

func (dao ProductDao) ReadAllProducts() ([]*dbModels.Product, error) {
	var products []*dbModels.Product
	query := "Select * from product"

	rows, err := prepareAndExecuteSelectQuery(dao.DB, query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var product dbModels.Product
		scanErr := rows.Scan(&product.Id, &product.BrandId, &product.CategoryId,
			&product.Name, &product.LongDescription,
			&product.ShortDescription, &product.Cost, &product.CardImage)
		if scanErr == nil {
			products = append(products, &product)
		}
	}
	return products, err
}
