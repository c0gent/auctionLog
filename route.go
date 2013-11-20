package main

import (
	"github.com/nsan1129/unframed"
)

var Router *unframed.Router

func route() {
	Router = unframed.NewRouter()
	r := Router

	r.Get("/sales/list", listSales)
	r.Get("/sale/form/{Id}", formSale)
	r.Post("/sale/save", saveSale)
	r.Post("/sale/delete", deleteSale)
	r.Dir("assets/")
	r.Dir("public/")

	r.Get("/", getHome)

	r.Serve("80")
}
