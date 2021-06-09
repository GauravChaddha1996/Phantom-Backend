package databasDaos

import (
	"database/sql"
	"phantom/dataLayer/dbModels"
)

type PropertyDao struct {
	DB *sql.DB
}

func (dao PropertyDao) CreateProperty(property dbModels.Property) (*int64, error) {
	query := "Insert into property (id, name) values (?, ?)"
	lastInsertId, err := prepareAndExecuteInsertQuery(dao.DB, query, property.Id, property.Name)
	if err != nil {
		return nil, err
	}
	return lastInsertId, nil
}

func (dao PropertyDao) ReadPropertyName(id int64) (*string, error) {
	var name *string
	query := "Select name from property where id = ? limit 1"

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

func (dao PropertyDao) ReadAllProperty() (*[]dbModels.Property, error) {
	query := "Select * from property"

	rows, err := prepareAndExecuteSelectQuery(dao.DB, query)
	if err != nil {
		return nil, err
	}

	var propertyArr []dbModels.Property

	for rows.Next() {
		var property dbModels.Property
		err = rows.Scan(&property.Id, &property.Name)
		if err != nil {
			return nil, err
		}
		propertyArr = append(propertyArr, property)
	}
	return &propertyArr, err
}
