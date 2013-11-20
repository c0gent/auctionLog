package main

import (
	"database/sql"
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

func (s *salesAdapter) list() []*sale {

	rows := s.Query(DB.Stmts["listSales"], 20)
	defer rows.Close()

	for rows.Next() {
		sa := s.newSale()
		s.ScanRows(rows, &sa.Id, &sa.ItemName, &sa.SoldPrice, &sa.Quality, &sa.Qty, &sa.ItemId, &sa.Comment)
	}
	return s.Sales
}

func (s *salesAdapter) crazyList() []*sale {
	scanner := func(rows *sql.Rows) {
		sa := s.newSale()
		s.ScanRows(rows, &sa.Id, &sa.ItemName, &sa.SoldPrice, &sa.Quality, &sa.Qty, &sa.ItemId, &sa.Comment)
	}
	s.CrazyQueryScan(scanner, DB.Stmts["listSales"], 20)
	return s.Sales
}

func (s *salesAdapter) show(id int) *salesAdapter {
	row := s.QueryRow(DB.Stmts["showSale"], id)

	sa := s.newSale()
	s.ScanRow(row, &sa.Id, &sa.ItemName, &sa.SoldPrice, &sa.Quality, &sa.Qty, &sa.ItemId, &sa.Comment)
	return s
}

func (s *salesAdapter) delete(id int) {
	s.Exec(DB.Stmts["deleteSale"], id)
}

func (s *salesAdapter) save(sa *sale) {
	if sa.Id == 0 {
		s.Exec(DB.Stmts["createSale"], sa.ItemName, sa.SoldPrice, sa.Quality, sa.Qty, sa.ItemId, sa.Comment)
	} else {
		s.Exec(DB.Stmts["updateSale"], sa.Id, sa.ItemName, sa.SoldPrice, sa.Quality, sa.Qty, sa.ItemId, sa.Comment)
	}
}

func (s *salesAdapter) newSale() *sale {
	sa := new(sale)
	s.Sales = append(s.Sales, sa)
	return sa
}
