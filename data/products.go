package data

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"time"

	"github.com/go-playground/validator"
)

type Product struct {
	ID          int		`json:"id"`
	Name        string	`json:"name" validate:"required"`
	Description string	`json:"description"`
	Price       float32	`json:"price" validate:"gt=0"`
	SKU         string	`json:"sku" validate:"required,sku"`
	CreatedOn   string	`json:"-"`
	UpdatedOn   string	`json:"-"`
	DeletedOn   string	`json:"-"`
}

type Products []*Product

func (p *Product) Validate() error {
	validate := validator.New()
	validate.RegisterValidation("sku", validateSku)
	return validate.Struct(p)
}

func validateSku(fl validator.FieldLevel) bool {
	re := regexp.MustCompile(`[a-z]+-[a-z]+-[a-z]+`)
	matches := re.FindAllString(fl.Field().String(), -1)

	return len(matches) == 1
}

func (p *Product) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)
}

func (p *Products) ToJSON(rw http.ResponseWriter) error {
	e := json.NewEncoder(rw)
	return e.Encode(p)
}

func GetProducts() Products {
	return productList
}

func AddProduct(p *Product) {
	p.ID = getNextID()
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

var ErrProductNotFound = fmt.Errorf("Product not found")

func findProduct(id int) (*Product, int, error) {
	for i, p := range productList {
		if p.ID == id {
			return p, i, nil
		}
	}

	return nil, -1, ErrProductNotFound
}


func getNextID() int {
	lp := productList[len(productList) - 1]
	return lp.ID + 1
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