package data

import (
	"encoding/json"
	"io"
	"time"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	SKU         string  `json:"sku"`
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"deletedOn,omitempty"`
}

type Products []*Product

func (p Products) ToJson(w io.Writer) error {
	return json.NewEncoder(w).Encode(p)
}

func GetProducts() Products {
	return productList
}

var productList = []*Product{
	&Product{
		ID:          1,
		Name:        "Latte",
		Description: "Frothy milky coffee",
		Price:       2.45,
		SKU:         "latte123",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          2,
		Name:        "Espresso",
		Description: "Short and strong coffee without milk",
		Price:       3.45,
		SKU:         "espreeso123",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          3,
		Name:        "Espresso",
		Description: "Short and strong coffee without milk",
		Price:       3.45,
		SKU:         "espreeso123",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          4,
		Name:        "Espresso",
		Description: "Short and strong coffee without milk",
		Price:       3.45,
		SKU:         "espreeso123",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          5,
		Name:        "Espresso",
		Description: "Short and strong coffee without milk",
		Price:       3.45,
		SKU:         "espreeso123",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          6,
		Name:        "Espresso",
		Description: "Short and strong coffee without milk",
		Price:       3.45,
		SKU:         "espreeso123",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          7,
		Name:        "Espresso",
		Description: "Short and strong coffee without milk",
		Price:       3.45,
		SKU:         "espreeso123",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          8,
		Name:        "Espresso",
		Description: "Short and strong coffee without milk",
		Price:       3.45,
		SKU:         "espreeso123",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          9,
		Name:        "Espresso",
		Description: "Short and strong coffee without milk",
		Price:       3.45,
		SKU:         "espreeso123",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          10,
		Name:        "Espresso",
		Description: "Short and strong coffee without milk",
		Price:       3.45,
		SKU:         "espreeso123",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}
