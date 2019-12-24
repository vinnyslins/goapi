package models

import "goapi/db"

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

	productsQuery, err := db.Query("SELECT * FROM products")

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
