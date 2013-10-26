package mdl

import (
	"fmt"
)

type Sale struct {
	Id        int
	ItemName  string
	SoldPrice int
	Quality   int
	Qty       int
	ItemId    int
	Comment   string
}

func ListSales() []*Sale {

	rows, err := DB.stmts.listSales.Query()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	defer rows.Close()

	var sales []*Sale

	for rows.Next() {
		s := new(Sale)
		err := rows.Scan(&s.Id, &s.ItemName, &s.SoldPrice, &s.Quality, &s.Qty, &s.ItemId, &s.Comment)
		if err != nil {
			fmt.Println(err)
		}
		sales = append(sales, s)
	}

	return sales
}
