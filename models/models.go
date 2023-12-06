package models

import (
	"fmt"
	"loja/db"
	"strconv"
)

type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func ListProducts() []Produto {
	db := db.ConnectDatabase()
	defer db.Close()

	allProducts, err := db.Query("select * from produtos")

	if err != nil {
		panic(err.Error())
	}

	produto := Produto{}
	produtos := []Produto{}

	for allProducts.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = allProducts.Scan(&id, &nome, &descricao, &preco, &quantidade)

		if err != nil {
			panic(err.Error())
		}

		produto.Id = id
		produto.Nome = nome
		produto.Descricao = descricao
		produto.Quantidade = quantidade
		produto.Preco = preco

		produtos = append(produtos, produto)

	}

	return produtos
}

func GetProduct(id string) Produto {
	db := db.ConnectDatabase()
	defer db.Close()

	productDb, err := db.Query("select * from produtos where id=$1", id)

	if err != nil {
		panic(err.Error())
	}

	product := Produto{}

	for productDb.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = productDb.Scan(&id, &nome, &descricao, &preco, &quantidade)

		if err != nil {
			panic(err.Error())
		}

		product.Id = id
		product.Nome = nome
		product.Descricao = descricao
		product.Preco = preco
		product.Quantidade = quantidade

	}
	return product
}

func CreateProduct(nome, descricao string, preco float64, quantidade int) {
	db := db.ConnectDatabase()
	defer db.Close()

	create, err := db.Prepare("INSERT INTO produtos (nome, descricao, preco, quantidade) VALUES ($1, $2, $3, $4);")

	if err != nil {
		panic(err.Error())
	}

	create.Exec(nome, descricao, preco, quantidade)

}

func EditProduct(id string, nome string, descricao string, preco float64, quantidade int) error {
	db := db.ConnectDatabase()
	defer db.Close()

	productId, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println(err)
	}

	edit, err := db.Prepare("UPDATE produtos SET nome=$1, descricao=$2, preco=$3, quantidade=$4 WHERE id=$5;")
	if err != nil {
		fmt.Println(err)
	}

	_, err = edit.Exec(nome, descricao, preco, quantidade, productId)
	return err
}

func DeleteProduct(id string) {
	db := db.ConnectDatabase()
	defer db.Close()

	deleteProduct, err := db.Prepare("delete from produtos where id=$1")

	if err != nil {
		panic(err.Error())
	}

	deleteProduct.Exec(id)

}
