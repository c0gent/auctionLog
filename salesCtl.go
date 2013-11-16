package main

import (
	"github.com/nsan1129/auctionLog/log"
	"net/http"
)

func listSales(w http.ResponseWriter, r *http.Request) {

	sm := new(salesMdl).getList()

	err := templates.ExecuteTemplate(w, "listSales", sm)
	if err != nil {
		log.Error(err)
	}
}

func composeSale(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, "composeSale", "None Yet")
	if err != nil {
		log.Error(err)
	}
}

func createSale(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Error(err)
	}
	sm := new(salesMdl)

	err = DB.decoder.Decode(sm.newSale(), r.PostForm)
	if err != nil {
		log.Error(err)
	}
	sm.commit()

	http.Redirect(w, r, "/sales/list", http.StatusFound)

}
