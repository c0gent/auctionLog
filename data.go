package main

import (
	"database/sql"
	"encoding/gob"
	_ "github.com/bmizerany/pq"
	"github.com/gorilla/schema"
)

type dbHandle struct {
	*sql.DB
	cdd     dbDialogue
	stmts   *stmtStore
	decoder *schema.Decoder
}

func (d *dbHandle) init() *dbHandle {
	var err error
	d.DB, err = sql.Open("postgres", "user=postgres password=postgres dbname=auction_log_dev sslmode=disable")
	if err != nil {
		panic(err)
	}

	d.cdd = CURRENTDBDIALOGUE

	return d
}

func initDB() {

	DB = new(dbHandle).init()
	DB.stmts = new(stmtStore).init()
	DB.decoder = schema.NewDecoder()

	gob.Register(&sale{})

}
