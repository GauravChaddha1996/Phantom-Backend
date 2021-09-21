package databaseDaos

import (
	"database/sql"
	"github.com/hashicorp/go-multierror"
	"phantom/dataLayer/dbModels"
)

type PropertyValueSqlDao struct {
	DB *sql.DB
}

func (dao PropertyValueSqlDao) CreatePropertyValue(propertyValue dbModels.PropertyValue) (*int64, error) {
	query := "Insert into property_value (id, property_id, name) values (?, ?, ?)"
	lastInsertId, err := prepareAndExecuteInsertQuery(dao.DB, query, propertyValue.Id, propertyValue.PropertyId, propertyValue.Name)
	if err != nil {
		return nil, err
	}
	return lastInsertId, nil
}

func (dao PropertyValueSqlDao) ReadPropertyValueName(id int64) (*string, error) {
	propertyValue, err := dao.ReadPropertyValueComplete(id)
	if err != nil {
		return nil, err
	}
	return &propertyValue.Name, err
}

func (dao PropertyValueSqlDao) ReadPropertyValueComplete(id int64) (*dbModels.PropertyValue, error) {
	var propertyValue dbModels.PropertyValue
	query := "Select * from property_value where id = ? limit 1"

	row, err := prepareAndExecuteSingleRowSelectQuery(dao.DB, query, id)
	if err != nil {
		return nil, err
	}

	err = row.Scan(&propertyValue.Id, &propertyValue.PropertyId, &propertyValue.Name)
	if err != nil {
		return nil, err
	}
	return &propertyValue, nil
}

func (dao PropertyValueSqlDao) ReadAllPropertyValues() (*[]dbModels.PropertyValue, error) {
	var propertyValueArr []dbModels.PropertyValue
	query := "Select * from property_value"

	rows, err := prepareAndExecuteSelectQuery(dao.DB, query)
	if err != nil {
		return nil, err
	}

	var allRowScanErrs error
	for rows.Next() {
		var propertyValue dbModels.PropertyValue
		rowScanErr := rows.Scan(&propertyValue.Id, &propertyValue.PropertyId, &propertyValue.Name)
		if rowScanErr != nil {
			allRowScanErrs = multierror.Append(allRowScanErrs, rowScanErr)
			continue
		}
		propertyValueArr = append(propertyValueArr, propertyValue)
	}

	return &propertyValueArr, allRowScanErrs
}

func (dao PropertyValueSqlDao) ReadAllPropertyValuesOfProperty(propertyId int64) (*[]dbModels.PropertyValue, error) {
	var propertyValueArr []dbModels.PropertyValue
	query := "Select * from property_value where property_id = ?"

	rows, err := prepareAndExecuteSelectQuery(dao.DB, query, propertyId)
	if err != nil {
		return nil, err
	}

	var allRowScanErrs error
	for rows.Next() {
		var propertyValue dbModels.PropertyValue
		rowScanErr := rows.Scan(&propertyValue.Id, &propertyValue.PropertyId, &propertyValue.Name)
		if rowScanErr != nil {
			allRowScanErrs = multierror.Append(allRowScanErrs, rowScanErr)
			continue
		}
		propertyValueArr = append(propertyValueArr, propertyValue)
	}

	return &propertyValueArr, allRowScanErrs
}
