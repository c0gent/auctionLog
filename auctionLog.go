package main

import (
	"encoding/gob"
	_ "github.com/bmizerany/pq"
	"github.com/nsan1129/unframed"
)

var db *unframed.DbHandle
var net *unframed.NetHandle

func main() {

	db = unframed.NewDB("postgres", "user=postgres password=postgres dbname=auction_log_dev sslmode=disable")
	defer db.Close()

	net = unframed.NewNet()

	prepareStatements()

	loadTemplates()

	route() //Set up routes and begin serving
}

func loadTemplates() {
	net.LoadTemplates(
		"tmpl/_base.html.tmpl",
		"tmpl/sales/listSales.html.tmpl",
		"tmpl/sales/formSale.html.tmpl",
		"tmpl/home/home.html.tmpl",
	)
}

func route() {

	gob.Register(&sale{})

	r := net

	r.Get("/sales/list", listSales)
	r.Get("/sale/form/{Id}", formSale)
	r.Post("/sale/save", saveSale)
	r.Post("/sale/delete", deleteSale)
	r.Dir("assets/")
	r.Dir("public/")

	r.Get("/", getHome)

	r.Serve("80")
}

/* TO DO
//----------


//---------
*/
