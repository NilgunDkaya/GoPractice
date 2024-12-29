package project

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Product struct {
	Id          int     `json:"id"`
	ProductName string  `json:"productName"`
	CategoryId  int     `json:"categoryId"`
	UnitPrice   float64 `json:"unitPrice"`
}

type Category struct {
	Id           int    `json:"id"`
	CategoryName string `json:"categoryName"`
}

func GetAllProducts() ([]Product, error) {
	response, err := http.Get("http://localhost:3000/products")
	if err != nil {
		fmt.Println(err)
	}

	defer response.Body.Close()

	bodyBytes, _ := io.ReadAll(response.Body)

	var products []Product
	json.Unmarshal(bodyBytes, &products)
	return products, nil
}

func AddProduct() (Product, error) {
	//product := Product{Id: 4, ProductName: "Telefon", CategoryId: 1, UnitPrice: 85000.0}
	//product := Product{ProductName: "Mikrofon", CategoryId: 1, UnitPrice: 85000.0}
	product := Product{Id: 10, ProductName: "Laptop Yükseltici", CategoryId: 1, UnitPrice: 100.0}
	jsonProduct, _ := json.Marshal(product)
	response, err := http.Post("http://localhost:3000/products",
		"application/json;charset=utf-8",
		bytes.NewBuffer(jsonProduct))
	if err != nil {
		return Product{}, err
	}
	defer response.Body.Close()

	bodyBytes, _ := io.ReadAll(response.Body)
	var productResponse Product
	json.Unmarshal(bodyBytes, &productResponse)
	return productResponse, nil
}
