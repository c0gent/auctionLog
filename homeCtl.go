package main

import (
	"github.com/nsan1129/auctionLog/log"
	"net/http"
)

type myData struct {
	importantShit string
}

func getHome(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, "home", nil)
	if err != nil {
		log.Error(err)
	}
}
