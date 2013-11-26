package main

import (
	"net/http"
)

func staticReg() {
	ptStatic()
}

func ptStatic() {
	net.TemplateFiles(
		"tmpl/_base.html.tmpl",
		"tmpl/static/home.html.tmpl",
	)
}

type myData struct {
	importantShit string
}

func getHome(w http.ResponseWriter, r *http.Request) {
	net.ExeTmpl(w, "home", nil)
}
