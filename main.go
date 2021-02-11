package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"phantom/dataLayer/daos"
)

func main() {
	db := openDB()
	dao := daos.PropertyDao{DB: db}
	name, err := dao.ReadPropertyName(1)
	if err != nil {
		println(err.Error())
	} else {
		println(*name)
	}
}

func openDB() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/phantom")
	if err != nil {
		panic(err)
	}
	return db
}
