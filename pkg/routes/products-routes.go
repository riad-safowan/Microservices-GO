package routes

import (
	"log"
	"os"

	"github.com/gorilla/mux"
	"github.com/riad-safowan/GO_MICROSERVICES/pkg/controllers"
)

var RegisterProductRoutes = func(router *mux.Router) {
	l := log.New(os.Stdout, "rest-api", log.LstdFlags)

	ph := controllers.NewProducts(l)

	getRouter := router.Methods("GET").Subrouter()
	getRouter.HandleFunc("/products", ph.GetProducts) // call direct function instead of  serveHTTP
	getRouter.HandleFunc("/products/{id:[0-9]+}", ph.GetProductById)

	postRouter := router.Methods("POST").Subrouter()
	postRouter.HandleFunc("/products", ph.AddProduct)
	postRouter.Use(ph.MiddlewareProductValidation)

	putRouter := router.Methods("PUT").Subrouter()
	putRouter.HandleFunc("/products/{id:[0-9]+}", ph.UpdateProduct)
	putRouter.Use(ph.MiddlewareProductValidation)

	deleteRouter := router.Methods("DELETE").Subrouter()
	deleteRouter.HandleFunc("/products/{id:[0-9]+}", ph.DeleteProduct)
	deleteRouter.HandleFunc("/products", ph.DeleteAllProducts)
	
}
