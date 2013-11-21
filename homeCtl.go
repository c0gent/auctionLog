package main

import (
	"net/http"
)

type myData struct {
	importantShit string
}

func getHome(w http.ResponseWriter, r *http.Request) {
	net.ExeTmpl(w, "home", nil)
}
