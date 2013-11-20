package main

import (
	"github.com/nsan1129/unframed"
)

func prepareStatements() {
	d := unframed.Dbd.Pg
	DB.AddStatement("createSale",
		d,
		`INSERT INTO "Sales" (
			"ItemName", 
			"SoldPrice", 
			"Quality", 
			"Qty", 
			"ItemId", 
			"Comment"
		) VALUES ($1, $2, $3, $4, $5, $6);`,
	)

	DB.AddStatement("listSales",
		d,
		`SELECT 
			"Id",
			COALESCE("ItemName", ''), 
			"SoldPrice", 
			COALESCE("Quality", 0), 
			"Qty", 
			COALESCE("ItemId", 0), 
			COALESCE("Comment", '')
		FROM "Sales" 
		ORDER BY "Id" DESC
		LIMIT $1;`,
	)

	DB.AddStatement("showSale",
		d,
		`SELECT 
			"Id",
			COALESCE("ItemName", ''), 
			"SoldPrice", 
			COALESCE("Quality", 0), 
			"Qty", 
			COALESCE("ItemId", 0), 
			COALESCE("Comment", '')
		FROM "Sales"
		WHERE "Id" = $1;`,
	)

	DB.AddStatement("updateSale",
		d,
		`UPDATE "Sales" SET 
			"ItemName" = $2, 
			"SoldPrice" = $3, 
			"Quality" = $4, 
			"Qty" = $5, 
			"ItemId" = $6, 
			"Comment" = $7
		WHERE "Id" = $1;`,
	)

	DB.AddStatement("deleteSale",
		d,
		`DELETE FROM "Sales"
		WHERE "Id" = $1;`,
	)

	DB.PrepareStatements()
}
