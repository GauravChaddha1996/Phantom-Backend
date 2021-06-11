package databasDaos

import (
	"database/sql"
	"phantom/dataLayer/dbModels"
)

type ProductImageSqlDao struct {
	DB *sql.DB
}

func (dao ProductImageSqlDao) AddProductImage(productImage dbModels.ProductImage) (*dbModels.ProductImage, error) {
	query := "Insert into product_image (product_id, url) values (?, ?)"
	_, err := prepareAndExecuteInsertQuery(dao.DB, query, productImage.ProductId, productImage.Url)
	if err != nil {
		return nil, err
	}
	return &productImage, nil
}

func (dao ProductImageSqlDao) ReadProductImages(productId int64) ([]*dbModels.ProductImage, error) {
	var images []*dbModels.ProductImage
	query := "Select * from product_image where product_id = ?"

	rows, err := prepareAndExecuteSelectQuery(dao.DB, query, productId)
	if err != nil {
		return nil, err
	}
	defer closeRows(rows)

	for rows.Next() {
		var image dbModels.ProductImage
		err = rows.Scan(&image.ProductId, &image.Url)
		if err != nil {
			continue
		}
		images = append(images, &image)
	}
	return images, err
}
