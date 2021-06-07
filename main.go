package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"phantom/dataLayer/daos"
	"phantom/dataLayer/dbModels"
)

func main() {
	db := openDB()
	defer db.Close()
	dao := daos.ProductToPropertyDao{db}
	_, err := dao.CreateProductToPropertyMapping(dbModels.ProductToProperty{
		ProductId:  1,
		PropertyId: 1,
		ValueId:    6,
	})
	if err != nil {
		log.Fatal(err)
	}
	mappingArr, err := dao.ReadAllProductToPropertyMapping(1)
	if err != nil {
		log.Fatal(err)
	}
	println("Mapping")
	for _, mapping := range *mappingArr {
		log.Println(mapping)
	}
}

func openDB() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/phantom")
	if err != nil {
		panic(err)
	}
	return db
}
