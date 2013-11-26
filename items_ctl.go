package main

import (
	"net/http"
)

func itemsReg() {
	itemsTemplates()
	itemsRoutes()
	itemsStmts()

	net.RegType(&item{})
}

func itemsTemplates() {
	net.TemplateFiles(
		"tmpl/items/items_list.html.tmpl",
		"tmpl/items/items_form.html.tmpl",
	)
}

func itemsRoutes() {
	sr := net.Subrouter("/items")
	sr.Get("/list", itemsList)
	sr.Get("/form/{Id}", itemsForm)
	sr.Post("/save", itemsSave)
	sr.Post("/delete", itemsDelete)
}

func itemsList(w http.ResponseWriter, r *http.Request) {
	sm := new(itemsAdapter).list()

	net.ExeTmpl(w, "itemsList", sm)
}

func itemsForm(w http.ResponseWriter, r *http.Request) {

	id := net.QueryUrl(":Id", r)

	sa := new(itemsAdapter)

	if id == 0 {
		_ = sa.newItem()
		net.ExeTmpl(w, "itemsForm", sa)
	} else {
		sa.show(id)
		net.ExeTmpl(w, "itemsForm", sa)
	}
}

func itemsSave(w http.ResponseWriter, r *http.Request) {

	sa := new(item)

	net.DecodeForm(sa, r)

	new(itemsAdapter).save(sa)

	http.Redirect(w, r, "/items/list", http.StatusFound)

}

func itemsDelete(w http.ResponseWriter, r *http.Request) {

	var sa struct{ Id int }

	net.DecodeForm(&sa, r)

	ss := new(itemsAdapter)
	ss.delete(sa.Id)
}
