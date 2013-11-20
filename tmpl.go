package main

import (
	"github.com/nsan1129/auctionLog/log"
	"html/template"
	"net/http"
)

var templates *template.Template

func loadTemplates() {
	templates = template.Must(template.ParseFiles(
		"tmpl/_base.html.tmpl",
		"tmpl/sales/listSales.html.tmpl",
		"tmpl/sales/formSale.html.tmpl",
		"tmpl/home/home.html.tmpl",
	))
}

func exeTmpl(w http.ResponseWriter, templateName string, templateData interface{}) {
	err := templates.ExecuteTemplate(w, templateName, templateData)
	if err != nil {
		log.Error(err)
	}
}
