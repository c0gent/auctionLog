package main

import (
	"github.com/nsan1129/auctionLog/log"
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

type salesMdl struct {
	Sales []*sale
}

func (s *salesMdl) getList() []*sale {

	rows, err := DB.stmts.listSales.Query(20)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		sa := new(sale)
		err := rows.Scan(&sa.Id, &sa.ItemName, &sa.SoldPrice, &sa.Quality, &sa.Qty, &sa.ItemId, &sa.Comment)
		if err != nil {
			log.Error(err)
		}
		s.Sales = append(s.Sales, sa)
	}

	return s.Sales
}

func (s *salesMdl) commit() {
	for _, sa := range s.Sales {
		_, err := DB.stmts.createSale.Exec(sa.ItemName, sa.SoldPrice, sa.Quality, sa.Qty, sa.ItemId, sa.Comment)
		if err != nil {
			log.Error(err)
		}
	}
}

func (s *salesMdl) newSale() *sale {
	ns := new(sale)
	s.Sales = append(s.Sales, ns)
	return ns
}
