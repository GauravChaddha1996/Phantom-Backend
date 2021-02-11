package daos

import (
	"database/sql"
	"phantom/dataLayer/dbModels"
)

type PropertyDao struct {
	DB *sql.DB
}

func (dao PropertyDao) CreateProperty(property dbModels.Property) error {
	query := "Insert into property (id, name) values (?, ?)"
	err := prepareAndExecuteInsertQuery(dao.DB, query, property.Id, property.Name)
	if err != nil {
		return err
	}
	return nil
}

func (dao PropertyDao) ReadPropertyName(id int) (*string, error) {
	var name *string
	query := "Select name from property where id = ? limit 1"

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