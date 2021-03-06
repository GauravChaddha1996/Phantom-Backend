package databaseDaos

import (
	"database/sql"
	"github.com/hashicorp/go-multierror"
	"phantom/apis/apiCommons"
	"phantom/dataLayer/dbModels"
	"strings"
)

type PropertySqlDao struct {
	DB *sql.DB
}

func (dao PropertySqlDao) CreateProperty(property dbModels.Property) (*int64, error) {
	query := "Insert into property (id, name) values (?, ?)"
	lastInsertId, err := prepareAndExecuteInsertQuery(dao.DB, query, property.Id, property.Name)
	if err != nil {
		return nil, err
	}
	return lastInsertId, nil
}

func (dao PropertySqlDao) ReadPropertyComplete(id int64) (*dbModels.Property, error) {
	var property dbModels.Property
	query := "Select * from property where id = ? limit 1"
	row, err := prepareAndExecuteSingleRowSelectQuery(dao.DB, query, id)
	if err != nil {
		return nil, err
	}

	err = row.Scan(&property.Id, &property.Name)
	if err != nil {
		return nil, err
	}
	return &property, err
}

func (dao PropertySqlDao) ReadProperties(ids []int64) (*[]dbModels.Property, error) {
	query := "Select * from property"
	if len(ids) != 0 {
		query = query + " where id in (?" + strings.Repeat(", ?", len(ids)-1) + ")"
	}
	rows, err := prepareAndExecuteSelectQuery(dao.DB, query, apiCommons.ToVarargInterface(ids)...)
	if err != nil {
		return nil, err
	}

	var propertyArr []dbModels.Property

	var allRowScanErrs error
	for rows.Next() {
		var property dbModels.Property
		rowScanErr := rows.Scan(&property.Id, &property.Name)
		if rowScanErr != nil {
			allRowScanErrs = multierror.Append(allRowScanErrs, rowScanErr)
			continue
		}
		propertyArr = append(propertyArr, property)
	}
	return &propertyArr, allRowScanErrs
}

func (dao PropertySqlDao) ReadAllProperty() (*[]dbModels.Property, error) {
	return dao.ReadProperties([]int64{})
}
