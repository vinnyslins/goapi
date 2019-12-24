package controllers

import (
	"goapi/models"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	products := models.GetProducts()
	templates.ExecuteTemplate(w, "Index", products)
}

func New(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")
		description := r.FormValue("description")

		price, err := strconv.ParseFloat(r.FormValue("price"), 64)
		if err != nil {
			log.Println("Erro na conversão do preço:", err)
		}

		quantity, err := strconv.Atoi(r.FormValue("quantity"))
		if err != nil {
			log.Println("Erro na conversão da quantidade:", err)
		}

		models.CreateProduct(name, description, price, quantity)

		http.Redirect(w, r, "/", 301)
	}
}

func Delete(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	models.DeleteProduct(id)

	http.Redirect(w, r, "/", 301)
}
