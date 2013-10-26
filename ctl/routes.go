package ctl

import (
	"github.com/gorilla/pat"
	"net/http"
)

func InitRouter() {
	a := pat.New()

	a.Get("/sales", ListSales)
	a.Get("/", homeGetHandler)

	http.Handle("/", a)
}

func Serve() {
	InitRouter()
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
