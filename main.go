package main

import (
	_ "github.com/bmizerany/pq"
	"github.com/nsan1129/unframed"
)

var db *unframed.DbHandle
var net *unframed.NetHandle

var cfgFile string = "/etc/auction-log/config.json"

func main() {
	cfg := unframed.ReadConfig(cfgFile)

	db = unframed.NewDB(cfg.DbType, cfg.ConnStr)
	defer db.Close()

	net = unframed.NewNet()

	registerControllers()

	db.PrepareStatements()
	net.LoadTemplates()

	route(cfg.ListenPort)
}

func writeConfig() {
	unframed.WriteConfig(&unframed.Config{"postgres", "user=postgres password=postgres dbname=auction_log_dev sslmode=disable", "80"}, cfgFile)
}

func registerControllers() {
	itemsReg()
	salesReg()
	staticReg()
}

func route(port string) {

	net.Serve(port)
}

/* TO DO
//----------


//---------
*/
