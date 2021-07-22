package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rijkerd/go-microservices/handlers"
)

func main() {
	router := mux.NewRouter()

	router.Handle("/products", handlers.GetProductHandler()).Methods("GET")
	router.Handle("/products", handlers.CreateProductHandler()).Methods("POST")

	// Create new server and assign the router
	server := http.Server{
		Addr:    ":5000",
		Handler: router,
	}

	fmt.Println("Staring Product Catalog server on Port 9090")

	// Start Server on defined port/host.
	server.ListenAndServe()
}
