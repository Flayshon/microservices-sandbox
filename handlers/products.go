package handlers

import (
	"flayshon/micro/data"
	"log"
	"net/http"
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
	switch r.Method {
		case http.MethodGet:
			p.getProducts(rw, r)
		case http.MethodPost:
			p.addProduct(rw, r)
		case http.MethodPut:
			p.l.Println("PUT", r.URL.Path)

			reg := regexp.MustCompile(`/([0-9]+)`)
			g := reg.FindAllStringSubmatch(r.URL.Path, -1)

			if len(g) != 1 {
				p.l.Println("Invalid URI more than one id")
				http.Error(rw, "Invalid URI", http.StatusBadRequest)
				return
			}

			if len(g[0]) != 2 {
				p.l.Println("Invalid URI more than one capture group")
				http.Error(rw, "Invalid URI", http.StatusBadRequest)
				return
			}

			idString := g[0][1]
			id, err := strconv.Atoi(idString)
			if err != nil {
				p.l.Println("Invalid URI unable to convert to number", idString)
				http.Error(rw, "Invalid URI", http.StatusBadRequest)
				return
			}

			p.updateProducts(id, rw, r)
			rw.WriteHeader(http.StatusNoContent)
		default:
			rw.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (p *Products) getProducts(rw http.ResponseWriter, r *http.Request) {
	lp := data.GetProducts()
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to output in JSON format", http.StatusInternalServerError)
	}
}

func (p *Products) addProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST product")

	prod := &data.Product{}
	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to unmarshal JSON.", http.StatusBadRequest)
	}

	data.AddProduct(prod)
}

func (p *Products) updateProducts(id int, rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle PUT Product")

	prod := &data.Product{}

	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
	}

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