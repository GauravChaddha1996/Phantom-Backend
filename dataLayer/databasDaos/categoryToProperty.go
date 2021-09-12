package databasDaos

import (
	"database/sql"
	"github.com/hashicorp/go-multierror"
	"phantom/dataLayer/dbModels"
)

type CategoryToPropertySqlDao struct {
	DB *sql.DB
}

func (dao CategoryToPropertySqlDao) CreateCategoryToPropertyLink(categoryId int64, propertyId int64) error {
	query := "INSERT into category_to_property Values(?,?)"
	_, err := prepareAndExecuteInsertQuery(dao.DB, query, &categoryId, &propertyId)
	if err != nil {
		return err
	}
	return nil
}

func (dao CategoryToPropertySqlDao) ReadAllCategoryToPropertyMapping() (*[]dbModels.CategoryToProperty, error) {
	var dataArr []dbModels.CategoryToProperty
	query := "SELECT * from category_to_property"
	rows, err := prepareAndExecuteSelectQuery(dao.DB, query)
	if err != nil {
		return nil, err
	}
	var allRowScanErrs error
	for rows.Next() {
		var data dbModels.CategoryToProperty
		rowScanErr := rows.Scan(&data.CategoryId, &data.PropertyId)
		if rowScanErr != nil {
			allRowScanErrs = multierror.Append(allRowScanErrs, rowScanErr)
			continue
		}
		dataArr = append(dataArr, data)
	}
	return &dataArr, allRowScanErrs
}

func (dao CategoryToPropertySqlDao) ReadCategoryToPropertyMappingForCategoryId(categoryId int64) (*[]int64, error) {
	var propertyIdArr []int64
	query := "SELECT property_id from category_to_property where category_id = ?"
	rows, err := prepareAndExecuteSelectQuery(dao.DB, query, categoryId)
	if err != nil {
		return nil, err
	}
	var allRowScanErrs error
	for rows.Next() {
		var propertyId int64
		rowScanErr := rows.Scan(&propertyId)
		if rowScanErr != nil {
			allRowScanErrs = multierror.Append(allRowScanErrs, rowScanErr)
			continue
		}
		propertyIdArr = append(propertyIdArr, propertyId)
	}
	return &propertyIdArr, allRowScanErrs
}
