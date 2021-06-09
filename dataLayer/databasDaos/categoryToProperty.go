package databasDaos

import (
	"database/sql"
	"phantom/dataLayer/dbModels"
)

type CategoryToPropertyDao struct {
	DB *sql.DB
}

func (dao CategoryToPropertyDao) CreateCategoryToPropertyLink(categoryId int64, propertyId int64) error {
	query := "INSERT into category_to_property Values(?,?)"
	_, err := prepareAndExecuteInsertQuery(dao.DB, query, &categoryId, &propertyId)
	if err != nil {
		return err
	}
	return nil
}

func (dao CategoryToPropertyDao) ReadAllCategoryToPropertyMapping() (*[]dbModels.CategoryToProperty, error) {
	var dataArr []dbModels.CategoryToProperty
	query := "SELECT * from category_to_property"
	rows, err := prepareAndExecuteSelectQuery(dao.DB, query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var data dbModels.CategoryToProperty
		scanErr := rows.Scan(&data.CategoryId, &data.PropertyId)
		if scanErr != nil {
			return nil, scanErr
		}
		dataArr = append(dataArr, data)
	}
	return &dataArr, nil
}

func (dao CategoryToPropertyDao) ReadCategoryToPropertyMappingForCategoryId(categoryId int64) (*[]int64, error) {
	var propertyIdArr []int64
	query := "SELECT property_id from category_to_property where category_id = ?"
	rows, err := prepareAndExecuteSelectQuery(dao.DB, query, categoryId)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var propertyId int64
		scanErr := rows.Scan(&propertyId)
		if scanErr != nil {
			return nil, scanErr
		}
		propertyIdArr = append(propertyIdArr, propertyId)
	}
	return &propertyIdArr, nil
}
