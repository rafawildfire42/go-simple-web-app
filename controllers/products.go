package controllers

import (
	"fmt"
	"loja/models"
	"net/http"
	"strconv"
	"text/template"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	products := models.ListProducts()
	temp.ExecuteTemplate(w, "Index", products)
}

func GetProduct(w http.ResponseWriter, r *http.Request) {
	productId := r.URL.Query().Get("id")
	product := models.GetProduct(productId)
	temp.ExecuteTemplate(w, "Edit", product)
}

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")
		description := r.FormValue("description")
		price := r.FormValue("price")
		amount := r.FormValue("amount")

		convertedPrice, err := strconv.ParseFloat(price, 64)

		if err != nil {
			fmt.Println("Erro ao converter o preço para float64.")
		}

		convertedAmount, err := strconv.Atoi(amount)

		if err != nil {
			fmt.Println("Erro ao converter a amount para int.")
		}

		models.CreateProduct(name, description, convertedPrice, convertedAmount)

	}

	http.Redirect(w, r, "r", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	productId := r.URL.Query().Get("id")
	models.DeleteProduct(productId)
	http.Redirect(w, r, "/", 301)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		name := r.FormValue("name")
		description := r.FormValue("description")
		price := r.FormValue("price")
		amount := r.FormValue("amount")

		convertedPrice, err := strconv.ParseFloat(price, 64)

		if err != nil {
			fmt.Println("Erro ao converter o preço para float64.")
		}

		convertedAmount, err := strconv.Atoi(amount)

		if err != nil {
			fmt.Println("Erro ao converter a amount para int.")
		}

		models.EditProduct(id, name, description, convertedPrice, convertedAmount)

	}

	http.Redirect(w, r, "/", 301)
}
