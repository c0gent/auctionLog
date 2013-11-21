package main

import (
	"github.com/nsan1129/unframed"
	"net/http"
)

func listSales(w http.ResponseWriter, r *http.Request) {
	sm := new(salesAdapter).list()

	net.ExeTmpl(w, "listSales", sm)
}

func formSale(w http.ResponseWriter, r *http.Request) {

	id := unframed.QueryUrl(":Id", r)

	sa := new(salesAdapter)

	if id == 0 {
		_ = sa.newSale()
		net.ExeTmpl(w, "formSale", sa)
	} else {
		sa.show(id)
		net.ExeTmpl(w, "formSale", sa)
	}
}

func saveSale(w http.ResponseWriter, r *http.Request) {

	sa := new(sale)

	net.DecodeForm(sa, r)

	new(salesAdapter).save(sa)

	http.Redirect(w, r, "/sales/list", http.StatusFound)

}

func deleteSale(w http.ResponseWriter, r *http.Request) {

	var sa struct{ Id int }

	net.DecodeForm(&sa, r)

	ss := new(salesAdapter)
	ss.delete(sa.Id)
}
