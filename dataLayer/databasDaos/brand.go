package databasDaos

import (
	"database/sql"
	"phantom/dataLayer/dbModels"
)

type BrandSqlDao struct {
	DB *sql.DB
}

func (dao BrandSqlDao) CreateBrand(brand dbModels.Brand) (*int64, error) {
	query := "Insert into brand (id, name, description) values (?, ?, ?)"
	lastInsertId, err := prepareAndExecuteInsertQuery(dao.DB, query, brand.Id, brand.Name, brand.Description)
	if err != nil {
		return nil, err
	}
	return lastInsertId, nil
}

func (dao BrandSqlDao) ReadBrandName(id int64) (*string, error) {
	brand, err := dao.ReadBrandComplete(id)
	if err != nil {
		return nil, err
	}
	return &brand.Name, nil
}

func (dao BrandSqlDao) ReadBrandComplete(id int64) (*dbModels.Brand, error) {
	var brand dbModels.Brand
	query := "Select * from brand where id = ? limit 1"

	row, err := prepareAndExecuteSingleRowSelectQuery(dao.DB, query, id)
	if err != nil {
		return nil, err
	}

	err = row.Scan(&brand.Id, &brand.Name, &brand.Description)
	if err != nil {
		return nil, err
	}
	return &brand, nil
}
