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
	produtos := models.GetAllProducts()
	temp.ExecuteTemplate(w, "Index", produtos)
}

func GetProduct(w http.ResponseWriter, r *http.Request) {
	productId := r.URL.Query().Get("id")
	produto := models.GetProduct(productId)
	temp.ExecuteTemplate(w, "Edit", produto)
}

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		precoConvertido, err := strconv.ParseFloat(preco, 64)

		if err != nil {
			fmt.Println("Erro ao converter o preço para float64.")
		}

		quantidadeConvertida, err := strconv.Atoi(quantidade)

		if err != nil {
			fmt.Println("Erro ao converter a quantidade para int.")
		}

		models.CreateProduct(nome, descricao, precoConvertido, quantidadeConvertida)

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
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		precoConvertido, err := strconv.ParseFloat(preco, 64)

		if err != nil {
			fmt.Println("Erro ao converter o preço para float64.")
		}

		quantidadeConvertida, err := strconv.Atoi(quantidade)

		if err != nil {
			fmt.Println("Erro ao converter a quantidade para int.")
		}

		models.EditProduct(id, nome, descricao, precoConvertido, quantidadeConvertida)

	}

	http.Redirect(w, r, "/", 301)
}
