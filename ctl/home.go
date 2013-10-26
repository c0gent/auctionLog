package ctl

import (
	"fmt"
	"github.com/nsan1129/auctionLog/log"
	"html/template"
	"net/http"
)

type myData struct {
	ImportantShit string
}

func homeGetHandler(w http.ResponseWriter, r *http.Request) {
	md := &myData{ImportantShit: "Very Important Shit"}
	t, err := template.ParseFiles(
		"tmpl/main.html",
	)
	if err != nil {
		log.L(err)
	}
	t.Execute(w, md)
	fmt.Println("Just Hit Home Handler!")
}
