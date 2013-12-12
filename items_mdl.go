package main

import (
	"github.com/nsan1129/unframed"
	"time"
)

type item struct {
	Id           int
	Name         string
	QtySold      int
	HighestPrice int
	LowestPrice  int
	AvgPrice     int
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type itemsAdapter struct {
	unframed.DataAdapter
	Items []*item
}

func (a *itemsAdapter) list() *itemsAdapter {
	a.Query(a.newItem, db.Stmts["listItems"], 20)
	return a
}

func (a *itemsAdapter) listBox() *itemsAdapter {
	a.Query(a.newItem, db.Stmts["listItems"], 100)
	return a
}

func (a *itemsAdapter) newItem() (inf interface{}) {
	n := new(item)
	a.Items = append(a.Items, n)
	inf = n
	return
}

func (a *itemsAdapter) show(id int) *itemsAdapter {
	a.Query(a.newItem, db.Stmts["showItem"], id)
	return a
}

func (a *itemsAdapter) delete(id int) {
	a.Exec(db.Stmts["deleteItem"], id)
}

func (a *itemsAdapter) save(d *item) {
	if d.Id == 0 {
		a.Exec(db.Stmts["createItem"], d.Name, d.QtySold, d.HighestPrice, d.LowestPrice, d.AvgPrice)
	} else {
		a.Exec(db.Stmts["updateItem"], d.Id, d.Name, d.QtySold, d.HighestPrice, d.LowestPrice, d.AvgPrice, time.Now())
	}
}
