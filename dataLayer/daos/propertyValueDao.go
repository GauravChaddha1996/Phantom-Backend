package daos

import (
	"database/sql"
	"phantom/dataLayer/dbModels"
)

type PropertyValueDao struct {
	DB *sql.DB
}

func (dao PropertyValueDao) CreatePropertyValue(propertyValue dbModels.PropertyValue) (*int64, error) {
	query := "Insert into property_value (id, property_id, name) values (?, ?, ?)"
	lastInsertId, err := prepareAndExecuteInsertQuery(dao.DB, query, propertyValue.Id, propertyValue.PropertyId, propertyValue.Name)
	if err != nil {
		return nil, err
	}
	return lastInsertId, nil
}

func (dao PropertyValueDao) ReadPropertyValueName(id int64) (*string, error) {
	var name *string
	query := "Select name from property_value where id = ? limit 1"

	row, err := prepareAndExecuteSingleRowSelectQuery(dao.DB, query, id)
	if err != nil {
		return nil, err
	}

	err = row.Scan(&name)
	if err != nil {
		return nil, err
	}
	return name, err
}

func (dao PropertyValueDao) ReadPropertyValueComplete(id int) (*dbModels.PropertyValue, error) {
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
