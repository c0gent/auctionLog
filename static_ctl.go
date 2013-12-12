package main

import (
	"net/http"
)

func staticReg() {
	ptStatic()

	prStatic()
}

func ptStatic() {
	net.TemplateFiles(
		"tmpl/_base.html.tmpl",
		"tmpl/static/home.html.tmpl",
	)
}

func prStatic() {
	net.Dir("assets/")
	net.Dir("public/")

	net.Get("/", getHome)
}

type myData struct {
	importantShit string
}

func getHome(w http.ResponseWriter, r *http.Request) {
	net.ExeTmpl(w, "home", nil)
}
