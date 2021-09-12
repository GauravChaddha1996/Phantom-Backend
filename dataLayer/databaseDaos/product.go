package databaseDaos

import (
	"database/sql"
	"github.com/hashicorp/go-multierror"
	"phantom/dataLayer/dbModels"
)

type ProductSqlDao struct {
	DB *sql.DB
}

func (dao ProductSqlDao) CreateProduct(product dbModels.Product) (*int64, error) {
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

func (dao ProductSqlDao) ReadProduct(id int64) (*dbModels.Product, error) {
	var product dbModels.Product
	query := "Select * from product where id = ? limit 1"

	row, err := prepareAndExecuteSingleRowSelectQuery(dao.DB, query, id)
	if err != nil {
		return nil, err
	}

	err = row.Scan(&product.Id, &product.BrandId, &product.CategoryId,
		&product.Name, &product.LongDescription,
		&product.ShortDescription, &product.Cost, &product.CardImage, &product.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &product, err
}

func (dao ProductSqlDao) ReadAllProducts() (*[]dbModels.Product, error) {
	var products []dbModels.Product
	query := "Select * from product"

	rows, err := prepareAndExecuteSelectQuery(dao.DB, query)
	if err != nil {
		return nil, err
	}

	var allRowScanErrs error
	for rows.Next() {
		var product dbModels.Product
		rowScanErr := rows.Scan(&product.Id, &product.BrandId, &product.CategoryId,
			&product.Name, &product.LongDescription,
			&product.ShortDescription, &product.Cost, &product.CardImage, &product.CreatedAt)
		if rowScanErr != nil {
			allRowScanErrs = multierror.Append(allRowScanErrs, rowScanErr)
			continue
		}
		products = append(products, product)
	}
	return &products, allRowScanErrs
}
