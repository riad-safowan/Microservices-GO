package handlers

import (
	"log"
	"net/http"
	"p1/data"
	"regexp"
	"strconv"
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

	} else if r.Method == http.MethodPost {

		p.addProduct(rw, r)

	} else if r.Method == http.MethodPut {

		re := regexp.MustCompile(`products/([0-9]+)`)
		g := re.FindAllStringSubmatch(r.URL.Path, -1)

		if len(g) != 1 {
			http.Error(rw, "Invalid URI 1", http.StatusBadRequest)
		} else {
			if len(g[0]) != 2 {
				http.Error(rw, "Invalid URI 2", http.StatusBadRequest)
			} else {
				idString := g[0][1]
				id, err := strconv.Atoi(idString)
				if err != nil {
					p.l.Println("stconv error")
				}else{
					p.updateProduct(id, rw, r)
				}
			}
		}


	} else {
		rw.WriteHeader(http.StatusMethodNotAllowed)
	}
}

// ServeHttp implements the go http.Handler interface
func (p *Products) getProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Product Page is called")

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

func (p *Products) addProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Add product is called")

	prod := &data.Product{}
	err := prod.FromJson(r.Body)

	if err != nil {
		http.Error(rw, "unable to unmashal json", http.StatusBadRequest)
	} else {
		data.AddProduct(prod)
		p.getProducts(rw, r)
	}

}

func (p *Products) updateProduct(id int, rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Update product is called")

	prod := &data.Product{}
	err := prod.FromJson(r.Body)

	if err != nil {
		http.Error(rw, "unable to unmashal json", http.StatusBadRequest)
	} else {
		err = data.UpdateProduct(id, prod)
		
		if err == data.ErrProductNotFound {
			http.Error(rw, "Product not found", http.StatusNotFound)
			return
		}
	
		if err != nil {
			http.Error(rw, "Product not found", http.StatusInternalServerError)
			return
		}
	}

}
