package handlers

import (
	"log"
	"net/http"
	"flayshon/micro/data"
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
		default:
			rw.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (ph *Products) getProducts(rw http.ResponseWriter, r *http.Request) {
	lp := data.GetProducts()
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to output in JSON format", http.StatusInternalServerError)
	}
}
