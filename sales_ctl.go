package main

import (
	"net/http"
)

func salesReg() {
	salesTemplates()
	salesRoutes()
	salesStmts()

	net.RegType(&item{})
}

func salesTemplates() {
	net.TemplateFiles(
		"tmpl/sales/sales_list.html.tmpl",
		"tmpl/sales/sales_form.html.tmpl",
	)
}

func salesRoutes() {
	sr := net.Subrouter("/sales")
	sr.Get("/list", salesList)
	sr.Get("/form/{Id}", salesForm)
	sr.Post("/save", salesSave)
	sr.Post("/delete", salesDelete)
}

func salesList(w http.ResponseWriter, r *http.Request) {
	sm := new(salesAdapter).list()

	net.ExeTmpl(w, "salesList", sm)
}

func salesForm(w http.ResponseWriter, r *http.Request) {

	id := net.QueryUrl(":Id", r)

	sa := new(salesAdapter)

	if id == 0 {
		_ = sa.newSale()
		net.ExeTmpl(w, "salesForm", sa)
	} else {
		sa.show(id)
		net.ExeTmpl(w, "salesForm", sa)
	}
}

func salesSave(w http.ResponseWriter, r *http.Request) {

	sa := new(sale)

	net.DecodeForm(sa, r)

	new(salesAdapter).save(sa)

	http.Redirect(w, r, "/sales/list", http.StatusFound)

}

func salesDelete(w http.ResponseWriter, r *http.Request) {

	var sa struct{ Id int }

	net.DecodeForm(&sa, r)

	ss := new(salesAdapter)
	ss.delete(sa.Id)
}
