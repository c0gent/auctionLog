package main

import (
	"html/template"
)

var templates *template.Template

func loadTemplates() {
	templates = template.Must(template.ParseFiles(
		"tmpl/_base.html.tmpl",
		"tmpl/sales/listSales.html.tmpl",
		"tmpl/sales/composeSale.html.tmpl",
		"tmpl/home/home.html.tmpl",
	))
}
