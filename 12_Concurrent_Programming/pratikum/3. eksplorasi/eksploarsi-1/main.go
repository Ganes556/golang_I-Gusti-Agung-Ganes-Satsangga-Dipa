package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Product struct {
	Title    string  `json:"title"`
	Price    float64 `json:"price"`
	Category string  `json:"category"`
}

func Get(url string, ch chan<- []Product) {	// ch chan<- []Product artinya ch adalah channel yang hanya bisa menerima data
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	var products []Product
	err = json.Unmarshal(body, &products)
	if err != nil {
		log.Fatal(err)
	}
	ch <- products
}

func main() {
	ch := make(chan []Product)

	go Get("https://fakestoreapi.com/products", ch)
	
	fmt.Println("products data")
	for _, product := range <-ch {
		fmt.Println("===")
		fmt.Printf("title: %s\n", product.Title)
		fmt.Printf("price: %f\n", product.Price)
		fmt.Printf("category: %s\n", product.Category)
		fmt.Println("===")
	}
}
