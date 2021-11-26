package controllers

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/riad-safowan/GO_MICROSERVICES/pkg/models"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

// ServeHttp implements the go http.Handler interface
func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Product Page is called")
	// lp := models.GetProducts()
	lp := models.GetAllProducts()

	err := lp.ToJson(rw)
	if err != nil {
		http.Error(rw, "unable to marshal json", http.StatusInternalServerError)
		return
	}
	// fmt.Fprint(rw, "Your IP: "+ReadUserIP(r))
}

func (p *Products) GetProductById(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Product Page is called")
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		p.l.Println("Unable to convert id")
		http.Error(rw, "Unable to convert id", http.StatusBadRequest)
		return
	}

	sp, _, err := models.GetProductById(id)

	if err == models.ErrProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(rw, "Product not found", http.StatusInternalServerError)
		return
	}

	err = sp.ToJson(rw)
	if err != nil {
		p.l.Println("unable to marshal json")
		http.Error(rw, "unable to marshal json", http.StatusInternalServerError)
	}
}

func (p *Products) AddProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Add product is called")
	prod := r.Context().Value(KeyProduct{}).(*models.Product)
	// models.AddProduct(prod)
	models.CreateProduct(prod)
	p.GetProducts(rw, r)
}

func (p *Products) UpdateProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Update product is called")
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "Unable to convert id", http.StatusBadRequest)
		return
	}
	prod := r.Context().Value(KeyProduct{}).(*models.Product)
	err = models.UpdateProduct(id, prod)
	if err == models.ErrProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(rw, "Product not found", http.StatusInternalServerError)
		return
	}

	p.GetProducts(rw, r)
}

func (p *Products) DeleteProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("DELETE product is called")
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "Unable to convert id", http.StatusBadRequest)
		return
	}
	err = models.DeleteProduct(id)
	if err == models.ErrProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(rw, "Product not found", http.StatusInternalServerError)
		return
	}

	p.GetProducts(rw, r)
}

func (p *Products) DeleteAllProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("DELETE product is called")
	models.DeleteAllProducts()
	p.GetProducts(rw, r)
}

type KeyProduct struct{}

func (p Products) MiddlewareProductValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		prod := &models.Product{}
		err := prod.FromJson(r.Body)
		if err != nil {
			http.Error(rw, "unable to unmashal json", http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
		req := r.WithContext(ctx)

		next.ServeHTTP(rw, req)
	})
}

func ReadUserIP(r *http.Request) string {
	IPAddress := r.Header.Get("X-Real-Ip")
	if IPAddress != "" {
		return "(Real) " + IPAddress
	}
	IPAddress = r.Header.Get("X-Forwarded-For")
	if IPAddress != "" {
		return "(Forwarded) " + IPAddress
	}
	return "(Remote) " + r.RemoteAddr
}
