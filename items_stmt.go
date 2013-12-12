package main

import (
	"github.com/nsan1129/unframed"
)

func itemsStmts() {
	d := unframed.Dbd.Pg

	db.AddStatement("listItems",
		d,
		`SELECT 
			*
		FROM items
		ORDER BY id DESC
		LIMIT $1;`,
	)

	db.AddStatement("createItem",
		d,
		`INSERT INTO items (
			name,
			qty_sold,
			highest_price,
			lowest_price,
			avg_price,
			created_at,
			updated_at
		) VALUES ($1, $2, $3, $4, $5, now(), now());`,
	)

	db.AddStatement("showItem",
		d,
		`SELECT
			*
		FROM items
		WHERE id = $1;`,
	)

	db.AddStatement("updateItem",
		d,
		`UPDATE items SET 
			name = $2,
			qty_sold = $3,
			highest_price = $4,
			lowest_price = $5,
			avg_price = $6,
			updated_at = $7
		WHERE id = $1;`,
	)

	db.AddStatement("deleteItem",
		d,
		`DELETE FROM items
		WHERE id = $1;`,
	)
}
