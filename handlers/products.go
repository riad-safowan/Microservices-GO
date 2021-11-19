package handlers

import (
	"log"
	"net/http"
	"p1/data"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}
func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		p.getProducts(rw, r)
	} else {
		rw.WriteHeader(http.StatusMethodNotAllowed)
	}
}

// ServeHttp implements the go http.Handler interface
func (h *Products) getProducts(rw http.ResponseWriter, r *http.Request) {
	h.l.Println("Product Page is called")

	lp := data.GetProducts()
	// d, err := json.Marshal(lp)

	// if err != nil {
	// 	http.Error(rw, "unable to marshal json", http.StatusInternalServerError)
	// } else {
	// 	rw.Write(d)
	// }

	err := lp.ToJson(rw)
	if err != nil {
		http.Error(rw, "unable to marshal json", http.StatusInternalServerError)
	}
}
