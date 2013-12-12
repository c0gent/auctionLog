package main

import (
	"github.com/nsan1129/unframed"
)

func salesStmts() {
	d := unframed.Dbd.Pg

	db.AddStatement("createSale",
		d,
		`INSERT INTO sales (
			sold_price, 
			quality, 
			qty, 
			item_id, 
			comment
		) VALUES ($1, $2, $3, $4, $5);`,
	)

	db.AddStatement("listSales",
		d,
		`SELECT 
			s.*,
			i.name as item_name

		FROM sales s, items i
		WHERE i.id = s.item_id
		ORDER BY s.id DESC
		LIMIT $1;`,
	)

	db.AddStatement("showSale",
		d,
		`SELECT 
			s.*,
			i.name as item_name
		FROM sales s, items i
		WHERE i.id = s.item_id 
		AND s.id = $1;`,
	)

	db.AddStatement("updateSale",
		d,
		`UPDATE sales SET 
			sold_price = $2, 
			quality = $3, 
			qty = $4, 
			item_id = $5, 
			comment = $6,
			updated_at = now()
		WHERE id = $1;`,
	)

	db.AddStatement("deleteSale",
		d,
		`DELETE FROM sales
		WHERE id = $1;`,
	)
}
