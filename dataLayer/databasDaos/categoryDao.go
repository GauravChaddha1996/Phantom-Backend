package databasDaos

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
	category, err := dao.ReadCategoryComplete(id)
	if err != nil {
		return nil, err
	}
	return &category.Name, err
}

func (dao CategoryDao) ReadCategoryComplete(id int64) (*dbModels.Category, error) {
	var category dbModels.Category
	query := "Select * from category where id = ? limit 1"

	row, err := prepareAndExecuteSingleRowSelectQuery(dao.DB, query, id)
	if err != nil {
		return nil, err
	}

	err = row.Scan(&category.Id, &category.Name, &category.Description)
	if err != nil {
		return nil, err
	}
	return &category, nil
}

func (dao CategoryDao) ReadAllCategories() (*[]dbModels.Category, error) {
	var categoryArr []dbModels.Category
	query := "Select * from category"

	rows, err := prepareAndExecuteSelectQuery(dao.DB, query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var category dbModels.Category
		rowReadErr := rows.Scan(&category.Id, &category.Name, &category.Description)
		if rowReadErr != nil {
			return nil, rowReadErr
		}
		categoryArr = append(categoryArr, category)
	}
	return &categoryArr, nil
}
