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
