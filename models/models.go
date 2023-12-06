package models

import (
	"fmt"
	"loja/db"
	"strconv"
)

type Product struct {
	Id          int
	Name        string
	Description string
	Price       float64
	Amount      int
}

func ListProducts() []Product {
	db := db.ConnectDatabase()
	defer db.Close()

	allProducts, err := db.Query("select * from produtos")

	if err != nil {
		panic(err.Error())
	}

	product := Product{}
	products := []Product{}

	for allProducts.Next() {
		var id, amount int
		var name, description string
		var price float64

		err = allProducts.Scan(&id, &name, &description, &price, &amount)

		if err != nil {
			panic(err.Error())
		}

		product.Id = id
		product.Name = name
		product.Description = description
		product.Amount = amount
		product.Price = price

		products = append(products, product)

	}

	return products
}

func GetProduct(id string) Product {
	db := db.ConnectDatabase()
	defer db.Close()

	productDb, err := db.Query("select * from produtos where id=$1", id)

	if err != nil {
		panic(err.Error())
	}

	product := Product{}

	for productDb.Next() {
		var id, amount int
		var name, description string
		var price float64

		err = productDb.Scan(&id, &name, &description, &price, &amount)

		if err != nil {
			panic(err.Error())
		}

		product.Id = id
		product.Name = name
		product.Description = description
		product.Price = price
		product.Amount = amount

	}
	return product
}

func CreateProduct(name, description string, price float64, amount int) {
	db := db.ConnectDatabase()
	defer db.Close()

	create, err := db.Prepare("INSERT INTO produtos (nome, descricao, preco, quantidade) VALUES ($1, $2, $3, $4);")

	if err != nil {
		panic(err.Error())
	}

	create.Exec(name, description, price, amount)

}

func EditProduct(id string, name string, description string, price float64, amount int) error {
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

	_, err = edit.Exec(name, description, price, amount, productId)
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
