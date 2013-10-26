package main

import (
	"github.com/nsan1129/auctionLog/ctl"
	"github.com/nsan1129/auctionLog/mdl"
	//"github.com/nsan1129/auctionLog/tmpl"
)

func main() {

	mdl.Init()
	defer mdl.DB.Close()
	//tmpl.Init()
	ctl.Serve()
}

/* TO DO
Optimize templates so they're read from disk only once and yet called on demand
//----------
var index = template.Must(template.ParseFiles(
	"views/_sales_table.html",
	"views/sales.html",
))
//---------


*/
