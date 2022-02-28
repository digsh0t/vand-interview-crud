package main

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/wintltr/vand-interview-crud-project/handler"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	credentials := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "DELETE", "OPTIONS"})
	origins := handlers.AllowedOrigins([]string{"*"})
	
	router.HandleFunc("/login", handler.LoginHandler).Methods("POST")
	router.HandleFunc("/register", handler.RegisterHandler).Methods("POST")
	router.HandleFunc("/listall", handler.ListAllWebAppUser).Methods("GET")

	//Store routing
	router.HandleFunc("/store-management/stores", handler.AddStoreHandler).Methods("POST")
	router.HandleFunc("/store-management/stores", handler.UpdateStoreHandler).Methods("PUT")
	router.HandleFunc("/store-management/stores/{id}", handler.RemoveStoreHandler).Methods("DELETE")
	router.HandleFunc("/store-management/stores/{id}", handler.StoreDetailHandler).Methods("GET")
	router.HandleFunc("/store-management/stores/list/{page}", handler.ListStoreByPageHandler).Methods("GET")

	//Product routing
	router.HandleFunc("/product-management/products", handler.AddProductHandler).Methods("POST")
	router.HandleFunc("/product-management/products", handler.UpdateProductHandler).Methods("PUT")
	router.HandleFunc("/product-management/products/{id}", handler.RemoveProductHandler).Methods("DELETE")
	router.HandleFunc("/product-management/products/{id}", handler.ProductDetailHandler).Methods("GET")
	router.HandleFunc("/product-management/products/list/{page}", handler.ListProductByPageHandler).Methods("GET")

	//User routing
	router.HandleFunc("/user-management/users/{id}", handler.UserDetailHandler).Methods("GET")

	log.Fatal(http.ListenAndServe(":8081", handlers.CORS(credentials, methods, origins)(router)))
}