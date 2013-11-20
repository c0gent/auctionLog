package main

import (
	"github.com/nsan1129/auctionLog/log"
	"github.com/nsan1129/unframed"
	"net/http"
)

func listSales(w http.ResponseWriter, r *http.Request) {
	sm := new(salesAdapter).crazyList()

	exeTmpl(w, "listSales", sm)
}

func formSale(w http.ResponseWriter, r *http.Request) {
	id := unframed.Atoi(r.URL.Query().Get(":Id"))

	sa := new(salesAdapter)

	if id == 0 {
		_ = sa.newSale()
		exeTmpl(w, "formSale", sa)
	} else {
		sa.show(id)
		exeTmpl(w, "formSale", sa)
	}
}

func saveSale(w http.ResponseWriter, r *http.Request) {

	sa := new(sale)

	err := r.ParseForm()
	if err != nil {
		log.Error(err)
	}
	err = DB.Decoder.Decode(sa, r.PostForm)
	if err != nil {
		log.Error(err)
	}

	new(salesAdapter).save(sa)

	http.Redirect(w, r, "/sales/list", http.StatusFound)

}

func deleteSale(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Error(err)
	}

	var sa struct{ Id int }

	err = DB.Decoder.Decode(&sa, r.PostForm)
	if err != nil {
		log.Error(err)
	}

	ss := new(salesAdapter)
	ss.delete(sa.Id)
}
