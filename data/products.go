package data

import (
	"encoding/json"
	"net/http"
	"time"
)

type Product struct {
	ID          int
	Name        string
	Description string
	Price       float32
	SKU         string
	CreatedOn   string `json:"-"`
	UpdatedOn   string `json:"-"`
	DeletedOn   string `json:"-"`
}

type Products []*Product

func (p *Products) ToJSON(rw http.ResponseWriter) error {
	e := json.NewEncoder(rw)
	return e.Encode(p)
}

func GetProducts() Products {
	return productList
}

var productList = Products {
	&Product {
		ID: 1,
		Name: "Latte",
		Description: "Frothy milky coffee",
		Price: 2.45,
		SKU: "abc323",
		CreatedOn: time.Now().UTC().String(),
		UpdatedOn: time.Now().UTC().String(),
		DeletedOn: time.Now().UTC().String(),
	},
	&Product {
		ID: 2,
		Name: "Espresso",
		Description: "Short and strong coffe without milk",
		Price: 1.99,
		SKU: "fjd34",
		CreatedOn: time.Now().UTC().String(),
		UpdatedOn: time.Now().UTC().String(),
		DeletedOn: time.Now().UTC().String(),
	},
}