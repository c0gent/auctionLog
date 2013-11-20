package main

import (
	"encoding/gob"
	_ "github.com/bmizerany/pq"
	"github.com/nsan1129/unframed"
)

var DB *unframed.DbHandler

func main() {

	gob.Register(&sale{})

	DB = unframed.NewDB("postgres", "user=postgres password=postgres dbname=auction_log_dev sslmode=disable")
	defer DB.Close()

	prepareStatements()

	loadTemplates()

	route() //Set up routes and begin serving
}

/* TO DO
//----------


//---------
*/
