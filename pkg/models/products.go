package models

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-playground/validator"
	"github.com/riad-safowan/GO_MICROSERVICES/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description"`
	Price       float32 `json:"price" validate:"gt=0"`
	SKU         string  `json:"sku"`
	CreatedAt   string  `json:"-"`
	UpdatedAt   string  `json:"-"`
	DeletedAt   string  `json:"-"`
}

type Products []*Product

var ErrProductNotFound = fmt.Errorf("Product not found")
var productList = []*Product{}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Product{})
}

func CreateProduct(p *Product) *Product {
	db.Create(p)
	return p
}
func GetAllProducts() Products {
	db.Find(&productList)
	return productList
}
func GetProductById(Id int) (*Product, *gorm.DB, error) {
	var p Product
	db := db.Where("ID=?", Id).Find(&p)
	if p.ID != 0 {
		return &p, db, nil
	}
	return nil, db, ErrProductNotFound
}
func DeleteProduct(Id int) error{
	db.Where("ID=?", Id).Delete(Product{})
	return nil
}
func DeleteAllProducts(){
	db.Exec("DELETE FROM products")
}

func UpdateProduct(id int, p *Product) error {
	db.Where("ID=?", id).Delete(Product{})
	p.ID = id
	db.Create(p)
	return nil
}

func (p *Product) Validate() error {
	return validator.New().Struct(p)
}

func (p Products) ToJson(w io.Writer) error {
	return json.NewEncoder(w).Encode(p)
}
func (p Product) ToJson(w io.Writer) error {
	return json.NewEncoder(w).Encode(p)
}

func (p *Product) FromJson(r io.Reader) error {
	return json.NewDecoder(r).Decode(p)
}
