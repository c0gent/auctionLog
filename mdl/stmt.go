package mdl

import (
	//"reflect"

	"database/sql"
)

type stmtTexts map[dbDialogue]string

type stmtStore struct {
	listSales *sql.Stmt
}

func (s *stmtStore) init() *stmtStore {
	s.makeStmts()
	return s
}

func (s *stmtStore) makeStmts() {
	listSalesT := stmtTexts{
		dbd.pg: `SELECT 
			id,
			COALESCE(item_name, ''), 
			sold_price, 
			COALESCE(quality, 0), 
			qty, 
			COALESCE(item_id, 0), 
			COALESCE(comment, '')
		FROM sales 
		ORDER BY id ASC 
		LIMIT 50`,
		dbd.mondb: `mondo sucks (not really though)`,
	}
	s.storeStmt(listSalesT, &s.listSales)
}

func (s *stmtStore) storeStmt(sts stmtTexts, storeField **sql.Stmt) {
	var err error
	*storeField, err = DB.Prepare(sts[DB.cdd])

	if err != nil {
		panic(err)
	}

}
