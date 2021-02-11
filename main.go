package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"phantom/dataLayer/daos"
	"phantom/dataLayer/dbModels"
)

func main() {
	db := openDB()
	dao := daos.PropertyValueDao{DB: db}
	propertyValue := dbModels.PropertyValue{
		PropertyId: 2,
		Name:       "Property Value Z",
	}
	lastInsertId, err := dao.CreatePropertyValue(propertyValue)
	if err != nil {
		println(err.Error())
	} else {
		propertyValue.Id = *lastInsertId
		name, err := dao.ReadPropertyValueName(propertyValue.Id)
		if err!=nil {
			println(err)
		} else {
			println(*name)
		}
	}
}

func openDB() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/phantom")
	if err != nil {
		panic(err)
	}
	return db
}
