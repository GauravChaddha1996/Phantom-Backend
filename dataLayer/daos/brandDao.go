package daos

import (
	"database/sql"
	"phantom/dataLayer/dbModels"
)

type BrandDao struct {
	DB *sql.DB
}

func (dao BrandDao) CreateBrand(brand dbModels.Brand) (*int64, error) {
	query := "Insert into brand (id, name, description) values (?, ?, ?)"
	lastInsertId, err := prepareAndExecuteInsertQuery(dao.DB, query, brand.Id, brand.Name, brand.Description)
	if err != nil {
		return nil, err
	}
	return lastInsertId, nil
}

func (dao BrandDao) ReadBrandName(id int64) (*string, error) {
	var name *string
	query := "Select name from brand where id = ? limit 1"

	rows, err := prepareAndExecuteSelectQuery(dao.DB, query, id)
	if err != nil {
		return nil, err
	}

	rows.Next()
	err = rows.Scan(&name)
	if err != nil {
		return nil, err
	}
	return name, err
}

func (dao BrandDao) ReadBrandComplete(id int64) (*dbModels.Brand, error) {
	var brand dbModels.Brand
	query := "Select * from brand where id = ? limit 1"

	rows, err := prepareAndExecuteSelectQuery(dao.DB, query, id)
	if err != nil {
		return nil, err
	}

	rows.Next()
	err = rows.Scan(&brand.Id, &brand.Name, &brand.Description)
	if err != nil {
		return nil, err
	}
	return &brand, nil
}
