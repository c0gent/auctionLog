package main

import (
	"database/sql"
)

type stmtTexts map[dbDialogue]string

type stmtStore struct {
	listSales  *sql.Stmt
	createSale *sql.Stmt
}

func (s *stmtStore) init() *stmtStore {
	s.makeStmts()
	return s
}

func (s *stmtStore) makeStmts() {
	listSalesT := stmtTexts{
		dbdPg: `SELECT 
			"Id",
			COALESCE("ItemName", ''), 
			"SoldPrice", 
			COALESCE("Quality", 0), 
			"Qty", 
			COALESCE("ItemId", 0), 
			COALESCE("Comment", '')
		FROM "Sales" 
		ORDER BY "Id" ASC 
		LIMIT $1`,
	}
	s.storeStmt(listSalesT, &s.listSales)

	createSaleT := stmtTexts{
		dbdPg: `insert into "Sales"
		("ItemName", "SoldPrice", "Quality", "Qty", "ItemId", "Comment") 
		values ($1, $2, $3, $4, $5, $6)`,
	}
	s.storeStmt(createSaleT, &s.createSale)
}

func (s *stmtStore) storeStmt(sts stmtTexts, storeField **sql.Stmt) {
	var err error
	*storeField, err = DB.Prepare(sts[DB.cdd])

	if err != nil {
		panic(err)
	}

}
