package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/rijkerd/go-microservices/entity"
)

// Get Product from data.json file
func GetProductHandler() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		data, err := entity.GetProducts()
		// Check if error exists
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
		}

		rw.Header().Add("content-type", "application/json")
		rw.WriteHeader(http.StatusFound)
		rw.Write(data)
	}
}

// CreateProductHandler is used to create new Product
func CreateProductHandler() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		data, err := ioutil.ReadAll(r.Body)

		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			return
		}

		var product entity.Product

		err = json.Unmarshal(data, &product)

		if err != nil {
			rw.WriteHeader(http.StatusExpectationFailed)
			rw.Write([]byte("Invalid Data Format"))
			return
		}

		err = entity.AddProduct(product)

		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}

		rw.WriteHeader(http.StatusCreated)
		rw.Write([]byte("Added New Product"))
	}
}
