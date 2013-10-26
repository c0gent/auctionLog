package mdl

import (
	"database/sql"
	"fmt"
	_ "github.com/bmizerany/pq"
)

var dbd dbDialogues = dbDialogues{pg: 1, mysql: 2, mondb: 3, sqlit: 4, nosql: 5, other: 6}

var DB *DbHandle

//This will be read from a text file eventually
var CURRENTDBDIALOGUE dbDialogue = dbd.pg

// Database Dialogues
type dbDialogue int
type dbDialogues struct {
	pg,
	mysql,
	mondb,
	sqlit,
	nosql,
	other dbDialogue
}

type DbHandle struct {
	*sql.DB
	cdd   dbDialogue
	stmts *stmtStore
}

func (d *DbHandle) init() *DbHandle {
	var err error
	d.DB, err = sql.Open("postgres", "user=postgres password=postgres dbname=auction_log_dev sslmode=disable")
	if err != nil {
		fmt.Println(err, `-------------`)
		panic(err)
	}

	d.cdd = CURRENTDBDIALOGUE //set current database dialogue to above constant

	return d
}

func Init() {

	DB = new(DbHandle).init()
	DB.stmts = new(stmtStore).init()

}
