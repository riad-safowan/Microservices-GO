package data

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator"
	"io"
	"log"
	"time"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description"`
	Price       float32 `json:"price" validate:"gt=0"`
	SKU         string  `json:"sku"`
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"deletedOn,omitempty"`
}

func (p *Product) Validate() error {
	return validator.New().Struct(p)
}

type Products []*Product

func (p Products) ToJson(w io.Writer) error {
	return json.NewEncoder(w).Encode(p)
}
func (p Product) ToJson(w io.Writer) error {
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

func UpdateProduct(id int, p *Product) error {
	_, pos, err := findProduct(id)
	if err != nil {
		return err
	}

	p.ID = id
	productList[pos] = p

	return nil
}
func GetProductById(id int) (*Product, error) {
	p, _, e := findProduct(id)
	return p, e
}

func findProduct(id int) (*Product, int, error) {
	for i, p := range productList {
		if p.ID == id {
			log.Print("GOT")
			return p, i, nil
		}
	}
	return nil, -1, fmt.Errorf("PNF")
}

func getNextId() int {
	return productList[len(productList)-1].ID + 1
}

var ErrProductNotFound = fmt.Errorf("Product not found")

var productList = []*Product{
	{
		ID:          1,
		Name:        "Latte",
		Description: "Frothy milky coffee",
		Price:       2.45,
		SKU:         "latte123",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	{
		ID:          2,
		Name:        "Espresso",
		Description: "Short and strong coffee without milk",
		Price:       3.45,
		SKU:         "espreeso123",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	{
		ID:          3,
		Name:        "Espasedfasewfwaef eaf weresdfgvaersrdfawefso",
		Description: "Short anergavsdfaertgreaaed strong coffee wittgsergaergraefgashout milk",
		Price:       3.45,
		SKU:         "espreeso123",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	{
		ID:          4,
		Name:        "Espresefsadfadersgsrzgvbrafsso",
		Description: "Short and wertwerfsaeergfaewfastrong coffee without milk",
		Price:       3.45,
		SKU:         "espreeso123",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	{
		ID:          5,
		Name:        "Espresso",
		Description: "Short and strrertong coffee without milk",
		Price:       3.45,
		SKU:         "espreeso123",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	{
		ID:          6,
		Name:        "Espresreterterfgergdfreterso",
		Description: "Short andertgwergarefgergt strong coffee without milk",
		Price:       3.45,
		SKU:         "espreeso123",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	{
		ID:          7,
		Name:        "Espresdafsdafsadfsdfsdafsdfasdfsadfsso",
		Description: "Short asdfsdafsdafadsfand strong coffee without milk",
		Price:       3.45,
		SKU:         "espreefsadfasdfsdso123",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	{
		ID:          8,
		Name:        "Espressfasdfsdafasdfsdafasdo",
		Description: "Short fasdfsadfsdafsda strong coffee without milk",
		Price:       3.45,
		SKU:         "gsdgsdgdsgdfs",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	{
		ID:          9,
		Name:        "dfgadsfgsdfgdsfgdfs",
		Description: "Shasdfgasdfgasdgsdf",
		Price:       3.45,
		SKU:         "espreeso123",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	{
		ID:          10,
		Name:        "Espresso",
		Description: "Short and st-podjfgpiodufhbngpiaudfhgief oiudfgh sidofuhgpidfh gpaedrighj without milk",
		Price:       3.45,
		SKU:         "espreeso123",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}
