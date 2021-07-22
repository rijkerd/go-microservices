package handlers

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rijkerd/go-microservices/entity"
)

// Get Product from data.json file
func GetProductsHandler() http.HandlerFunc {
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

//
func GetProductHandler() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		// Read product ID
		productID := mux.Vars(r)["id"]
		product, err := entity.GetProduct(productID)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}
		responseData, err := json.Marshal(product)
		if err != nil {
			// Check if it is No product error or any other error
			if errors.Is(err, entity.ErrNoProduct) {
				// Write Header if no related product found.
				rw.WriteHeader(http.StatusNoContent)
			} else {
				rw.WriteHeader(http.StatusInternalServerError)
			}
			return
		}
		// Write body with found product
		rw.Header().Add("content-type", "application/json")
		rw.WriteHeader(http.StatusFound)
		rw.Write(responseData)
	}
}

// DeleteProductHandler deletes the product with given ID.
func DeleteProductHandler() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		// Read product ID
		productID := mux.Vars(r)["id"]
		err := entity.DeleteProduct(productID)
		if err != nil {
			// Check if it is No product error or any other error
			if errors.Is(err, entity.ErrNoProduct) {
				// Write Header if no related product found.
				rw.WriteHeader(http.StatusNoContent)
			} else {
				rw.WriteHeader(http.StatusInternalServerError)
			}
			return
		}
		// Write Header with Accepted Status (done operation)
		rw.WriteHeader(http.StatusAccepted)
	}
}

// UpdateProductHandler deletes the product with given ID.
func UpdateProductHandler() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		// Read product ID
		productID := mux.Vars(r)["id"]
		err := entity.DeleteProduct(productID)
		if err != nil {
			if errors.Is(err, entity.ErrNoProduct) {
				rw.WriteHeader(http.StatusNoContent)
			} else {
				rw.WriteHeader(http.StatusInternalServerError)
			}
			return
		}
		// Read incoming JSON from request body
		data, err := ioutil.ReadAll(r.Body)
		// If no body is associated return with StatusBadRequest
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			return
		}
		// Check if data is proper JSON (data validation)
		var product entity.Product
		err = json.Unmarshal(data, &product)
		if err != nil {
			rw.WriteHeader(http.StatusExpectationFailed)
			rw.Write([]byte("Invalid Data Format"))
			return
		}
		// Addproduct with the requested body
		err = entity.AddProduct(product)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}
		// Write Header if no related product found.
		rw.WriteHeader(http.StatusAccepted)
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
