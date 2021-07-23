package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rijkerd/go-microservices/handlers"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

func main() {
	router := mux.NewRouter()

	router.Use(loggingMiddleware)

	router.Handle("/products", handlers.GetProductsHandler()).Methods("GET")
	router.Handle("/products", handlers.CreateProductHandler()).Methods("POST")
	router.Handle("/products/{id}", handlers.GetProductHandler()).Methods("GET")
	router.Handle("/products/{id}", handlers.DeleteProductHandler()).Methods("DELETE")
	router.Handle("/products/{id}", handlers.UpdateProductHandler()).Methods("PUT")

	// Create new server and assign the router
	server := http.Server{
		Addr:    ":5000",
		Handler: router,
		// Handler: handlers.AuthHandler(router),
	}

	fmt.Println("Staring Product Catalog server on Port 5000")

	// Start Server on defined port/host.
	server.ListenAndServe()
}
