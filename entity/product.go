package entity

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
)

// Define struct for a product
type Product struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	IsAvailable bool    `json:"isAvailable"`
}

var ErrNoProduct = errors.New("No Product Found")

// Fetch Products
func GetProducts() ([]byte, error) {
	// Read from Json File
	data, err := ioutil.ReadFile("./data/data.json")
	if err != nil {
		return nil, err
	}

	return data, nil
}

// Add Products
func AddProduct(product Product) error {
	var products []Product

	data, err := ioutil.ReadFile("./data/data.json")

	if err != nil {
		return err
	}
	// Read Json File and add to Memory
	err = json.Unmarshal(data, &products)
	if err != nil {
		return err
	}

	products = append(products, product)

	// Write Update data
	updatedData, err := json.Marshal(products)

	if err != nil {
		return nil
	}

	err = ioutil.WriteFile("./data/data.json", updatedData, os.ModePerm)

	if err != nil {
		return err
	}

	return nil

}
