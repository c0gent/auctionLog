package main

import (
	"github.com/nsan1129/unframed"
)

type sale struct {
	Id        int
	ItemName  string
	SoldPrice int
	Quality   int
	Qty       int
	ItemId    int
	Comment   string
}

type salesAdapter struct {
	unframed.DataAdapter
	Sales []*sale
}

func (s *salesAdapter) list() *salesAdapter {
	s.Query(s.newSale, db.Stmts["listSales"], 20)
	return s
}

func (s *salesAdapter) show(id int) *salesAdapter {
	s.Query(s.newSale, db.Stmts["showSale"], id)
	return s
}

func (s *salesAdapter) delete(id int) {
	s.Exec(db.Stmts["deleteSale"], id)
}

func (s *salesAdapter) save(sa *sale) {
	if sa.Id == 0 {
		s.Exec(db.Stmts["createSale"], sa.ItemName, sa.SoldPrice, sa.Quality, sa.Qty, sa.ItemId, sa.Comment)
	} else {
		s.Exec(db.Stmts["updateSale"], sa.Id, sa.ItemName, sa.SoldPrice, sa.Quality, sa.Qty, sa.ItemId, sa.Comment)
	}
}

func (s *salesAdapter) newSale() (inf interface{}) {
	sa := new(sale)
	s.Sales = append(s.Sales, sa)
	inf = sa
	return
}
