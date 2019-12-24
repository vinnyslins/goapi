package models

import (
	"goapi/db"
)

type Product struct {
	ID          int
	Name        string
	Description string
	Price       float64
	Quantity    int
}

func GetProducts() []Product {
	db := db.Connect()
	defer db.Close()

	productsQuery, err := db.Query("SELECT * FROM products ORDER BY id ASC")

	if err != nil {
		panic(err.Error())
	}

	products := []Product{}

	for productsQuery.Next() {
		var id, quantity int
		var name, description string
		var price float64

		err := productsQuery.Scan(&id, &name, &description, &price, &quantity)

		if err != nil {
			panic(err.Error())
		}

		p := Product{id, name, description, price, quantity}

		products = append(products, p)
	}

	return products
}

func CreateProduct(name, description string, price float64, quantity int) {
	db := db.Connect()
	defer db.Close()

	productInsert, err := db.Prepare("INSERT INTO products(name, description, price, quantity) VALUES($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}

	productInsert.Exec(name, description, price, quantity)
}

func DeleteProduct(id string) {
	db := db.Connect()
	defer db.Close()

	productDelete, err := db.Prepare("DELETE FROM products WHERE id=$1")
	if err != nil {
		panic(err.Error())
	}

	productDelete.Exec(id)
}

func GetProduct(id string) Product {
	db := db.Connect()
	defer db.Close()

	productQuery, err := db.Query("SELECT * FROM products WHERE id=$1", id)
	if err != nil {
		panic(err.Error())
	}

	var product Product

	for productQuery.Next() {
		var name, description string
		var id, quantity int
		var price float64

		err = productQuery.Scan(&id, &name, &description, &price, &quantity)
		if err != nil {
			panic(err.Error())
		}

		product = Product{id, name, description, price, quantity}
	}

	return product
}

func UpdateProduct(id int, name, description string, price float64, quantity int) {
	db := db.Connect()
	defer db.Close()

	productUpdate, err := db.Prepare("UPDATE products SET name=$1, description=$2, price=$3, quantity=$4 WHERE id=$5")
	if err != nil {
		panic(err.Error())
	}

	productUpdate.Exec(name, description, price, quantity, id)
}
