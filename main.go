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
	dao := daos.ProductDao{db}
	productId, err := dao.CreateProduct(dbModels.Product{
		Id:               1,
		BrandId:          1,
		CategoryId:       1,
		Name:             "Test product 1 ",
		LongDescription:  "Prodyct 1 long desc",
		ShortDescription: "product 1 short desc",
		Cost:             100,
		CardImage:        "image_url",
	})
	if err != nil {
		log.Fatal(err)
	}
	product, err := dao.ReadProduct(*productId)
	if err != nil {
		log.Fatal(err)
	}
	log.Print(product)
	log.Print(product.Name)
}

func openDB() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/phantom")
	if err != nil {
		panic(err)
	}
	return db
}
