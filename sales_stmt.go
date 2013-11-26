package main

import (
	"github.com/nsan1129/unframed"
)

func salesStmts() {
	d := unframed.Dbd.Pg

	db.AddStatement("createSale",
		d,
		`INSERT INTO sales (
			item_name, 
			sold_price, 
			quality, 
			qty, 
			item_id, 
			comment
		) VALUES ($1, $2, $3, $4, $5, $6);`,
	)

	db.AddStatement("listSales",
		d,
		`SELECT 
			id,
			item_name, 
			sold_price, 
			COALESCE(quality, 0), 
			qty, 
			item_id, 
			COALESCE(comment, '')
		FROM sales 
		ORDER BY id DESC
		LIMIT $1;`,
	)

	db.AddStatement("showSale",
		d,
		`SELECT 
			id,
			COALESCE(item_name, ''), 
			sold_price, 
			COALESCE(quality, 0), 
			qty, 
			COALESCE(item_id, 0), 
			COALESCE(comment, '')
		FROM sales
		WHERE id = $1;`,
	)

	db.AddStatement("updateSale",
		d,
		`UPDATE sales SET 
			item_name = $2, 
			sold_price = $3, 
			quality = $4, 
			qty = $5, 
			item_id = $6, 
			comment = $7
		WHERE id = $1;`,
	)

	db.AddStatement("deleteSale",
		d,
		`DELETE FROM sales
		WHERE id = $1;`,
	)
}
