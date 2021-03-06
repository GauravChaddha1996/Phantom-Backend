package databaseDaos

import (
	"database/sql"
	"github.com/hashicorp/go-multierror"
	"phantom/apis/apiCommons"
	"phantom/dataLayer/dbModels"
	"strings"
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

func (dao BrandSqlDao) ReadBrands(ids []int64) (*[]dbModels.Brand, error) {
	var brands []dbModels.Brand
	query := "Select * from brand"
	if len(ids) != 0 {
		query = query + " where id in (?" + strings.Repeat(", ?", len(ids)-1) + ")"
	}
	rows, err := prepareAndExecuteSelectQuery(dao.DB, query, apiCommons.ToVarargInterface(ids)...)
	if err != nil {
		return nil, err
	}

	var allRowScanErrs error
	for rows.Next() {
		var brand dbModels.Brand
		rowScanErr := rows.Scan(&brand.Id, &brand.Name, &brand.Description)
		if rowScanErr != nil {
			allRowScanErrs = multierror.Append(allRowScanErrs, rowScanErr)
			continue
		}
		brands = append(brands, brand)
	}

	return &brands, allRowScanErrs
}

func (dao BrandSqlDao) ReadAllBrands() (*[]dbModels.Brand, error) {
	return dao.ReadBrands([]int64{})
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
