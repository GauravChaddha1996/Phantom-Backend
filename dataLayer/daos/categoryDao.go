package daos

import (
	"database/sql"
	"phantom/dataLayer/dbModels"
)

type CategoryDao struct {
	DB *sql.DB
}

func (dao CategoryDao) CreateCategory(category dbModels.Category) (*int64, error) {
	query := "Insert into category (id, name, description) values (?, ?, ?)"
	lastInsertId, err := prepareAndExecuteInsertQuery(dao.DB, query, category.Id, category.Name, category.Description)
	if err != nil {
		return nil, err
	}
	return lastInsertId, nil
}

func (dao CategoryDao) ReadCategoryName(id int64) (*string, error) {
	var name *string
	query := "Select name from category where id = ? limit 1"

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

func (dao CategoryDao) ReadCategoryComplete(id int64) (*dbModels.Category, error) {
	var category dbModels.Category
	query := "Select * from category where id = ? limit 1"

	rows, err := prepareAndExecuteSelectQuery(dao.DB, query, id)
	if err != nil {
		return nil, err
	}

	rows.Next()
	err = rows.Scan(&category.Id, &category.Name, &category.Description)
	if err != nil {
		return nil, err
	}
	return &category, nil
}
