package data

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
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

func (p *Product) FromJson(r io.Reader) error {
	return json.NewDecoder(r).Decode(p)
}

func GetProducts() Products {
	return productList
}

func AddProducts() {
	log.Print("New product added")
}

func AddProduct(p *Product) {
	p.ID = getNextId()
	productList = append(productList, p)
}

func UpdateProduct(id int, p*Product) error {
	_, pos, err := findProduct(id)
	if err != nil {
		return err
	}

	p.ID = id
	productList[pos] = p

	return nil
}

func findProduct(id int) (*Product, int, error) {
	for i, p := range productList {
		if p.ID == id {
			return p, i, nil
		}
	}

	return nil, -1, ErrProductNotFound 
}

func getNextId() int {
	return productList[len(productList)-1].ID + 1
}

var ErrProductNotFound = fmt.Errorf("Product not found")

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
		Name:        "Espasedfasewfwaef eaf weresdfgvaersrdfawefso",
		Description: "Short anergavsdfaertgreaaed strong coffee wittgsergaergraefgashout milk",
		Price:       3.45,
		SKU:         "espreeso123",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          4,
		Name:        "Espresefsadfadersgsrzgvbrafsso",
		Description: "Short and wertwerfsaeergfaewfastrong coffee without milk",
		Price:       3.45,
		SKU:         "espreeso123",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          5,
		Name:        "Espresso",
		Description: "Short and strrertong coffee without milk",
		Price:       3.45,
		SKU:         "espreeso123",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          6,
		Name:        "Espresreterterfgergdfreterso",
		Description: "Short andertgwergarefgergt strong coffee without milk",
		Price:       3.45,
		SKU:         "espreeso123",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          7,
		Name:        "Espresdafsdafsadfsdfsdafsdfasdfsadfsso",
		Description: "Short asdfsdafsdafadsfand strong coffee without milk",
		Price:       3.45,
		SKU:         "espreefsadfasdfsdso123",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          8,
		Name:        "Espressfasdfsdafasdfsdafasdo",
		Description: "Short fasdfsadfsdafsda strong coffee without milk",
		Price:       3.45,
		SKU:         "gsdgsdgdsgdfs",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          9,
		Name:        "dfgadsfgsdfgdsfgdfs",
		Description: "Shasdfgasdfgasdgsdf",
		Price:       3.45,
		SKU:         "espreeso123",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          10,
		Name:        "Espresso",
		Description: "Short and st-podjfgpiodufhbngpiaudfhgief oiudfgh sidofuhgpidfh gpaedrighj without milk",
		Price:       3.45,
		SKU:         "espreeso123",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}
