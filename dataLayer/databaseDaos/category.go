package databaseDaos

import (
	"database/sql"
	"github.com/hashicorp/go-multierror"
	"phantom/dataLayer/dbModels"
)

type CategorySqlDao struct {
	DB *sql.DB
}

func (dao CategorySqlDao) CreateCategory(category dbModels.Category) (*int64, error) {
	query := "Insert into category (id, name, description) values (?, ?, ?)"
	lastInsertId, err := prepareAndExecuteInsertQuery(dao.DB, query, category.Id, category.Name, category.Description)
	if err != nil {
		return nil, err
	}
	return lastInsertId, nil
}

func (dao CategorySqlDao) ReadCategoryName(id int64) (*string, error) {
	category, err := dao.ReadCategoryComplete(id)
	if err != nil {
		return nil, err
	}
	return &category.Name, err
}

func (dao CategorySqlDao) ReadCategoryComplete(id int64) (*dbModels.Category, error) {
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

func (dao CategorySqlDao) ReadAllCategories() (*[]dbModels.Category, error) {
	var categoryArr []dbModels.Category
	query := "Select * from category"

	rows, err := prepareAndExecuteSelectQuery(dao.DB, query)
	if err != nil {
		return nil, err
	}
	var allRowScanErrs error
	for rows.Next() {
		var category dbModels.Category
		rowScanErr := rows.Scan(&category.Id, &category.Name, &category.Description)
		if rowScanErr != nil {
			allRowScanErrs = multierror.Append(allRowScanErrs, rowScanErr)
			continue
		}
		categoryArr = append(categoryArr, category)
	}
	return &categoryArr, allRowScanErrs
}
