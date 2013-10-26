package ctl

import (
	"fmt"
	"github.com/nsan1129/auctionLog/mdl"
	"html/template"
	"net/http"
)

func ListSales(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles(
		"tmpl/sales/sales.html",
		"tmpl/sales/_sales_table.html",
	)
	if err != nil {
		fmt.Println(err)
	}

	sa := mdl.ListSales()

	err = t.Execute(w, sa)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Loaded with", len(sa), "records.")

}
