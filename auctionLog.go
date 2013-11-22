package main

import (
	_ "github.com/bmizerany/pq"
	"github.com/nsan1129/unframed"
)

var db *unframed.DbHandle
var net *unframed.NetHandle

func main() {

	db = unframed.NewDB("postgres", "user=postgres password=postgres dbname=auction_log_dev sslmode=disable")
	defer db.Close()

	net = unframed.NewNet()

	registerControllers()

	db.PrepareStatements()
	net.LoadTemplates()

	route() //Set up routes and begin serving
}

func registerControllers() {
	itemsReg()
	salesReg()
	staticReg()
}

func route() {

	net.Dir("assets/")
	net.Dir("public/")

	net.Get("/", getHome)

	net.Serve("80")
}

/* TO DO
//----------


//---------
*/
