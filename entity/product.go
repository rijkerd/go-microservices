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

// Get Product based on String
func GetProduct(id string) (Product, error) {
	data, err := ioutil.ReadFile("./data/data.json")
	if err != nil {
		return Product{}, err
	}

	var products []Product
	err = json.Unmarshal(data, &products)

	if err != nil {
		return Product{}, err
	}

	// iterate through product array
	for i := 0; i < len(products); i++ {
		// if we find one product with the given ID
		if products[i].ID == id {
			// return product
			return products[i], nil
		}
	}

	return Product{}, ErrNoProduct
}

// Delete Single Product
func DeleteProduct(id string) error {
	// Read JSON file
	data, err := ioutil.ReadFile("./data/data.json")
	if err != nil {
		return err
	}

	// read products
	var products []Product
	err = json.Unmarshal(data, &products)
	if err != nil {
		return err
	}

	for i := 0; i < len(products); i++ {
		// if we find one product with the given ID
		if products[i].ID == id {
			products = removeElement(products, i)
			// Write Updated JSON file
			updatedData, err := json.Marshal(products)
			if err != nil {
				return err
			}
			err = ioutil.WriteFile("./data/data.json", updatedData, os.ModePerm)
			if err != nil {
				return err
			}
			return nil
		}
	}
	return ErrNoProduct
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

// removeElement is used to remove element from product array at given index
func removeElement(arr []Product, index int) []Product {
	ret := make([]Product, 0)
	ret = append(ret, arr[:index]...)
	return append(ret, arr[index+1:]...)
}
