package main

import ()

var DB *dbHandle

func main() {

	initDB()
	defer DB.Close()
	loadTemplates()
	serve()
}

/* TO DO
//----------


//---------
*/

/*
--TERMINOLOGY--
-Actions (Must be Verbs)-
Go			Http		SQL				Purpose
------------------------------------------------------------
List		GET 		SELECT			display multiple records
Show		GET 		SELECT			display single record
Compose		GET			(none)			display composition controls/tools
Create		POST		INSERT			store new record(s)
Edit		GET			(none)			display editing controls/tools
Update		PUT			UPDATE			modify existing record(s)
Delete		DELETE		DELETE			destroy existing record(s)

NewXXX()			Return a new instance of something. Customary GO shorthand for GetNewXXX().

-Other Terms-
List(noun) = Table of Data, rows(multiple), etc.

*/

// Database Dialogues
type dbDialogue int

const (
	dbdPg    dbDialogue = 1001
	dbdMysql dbDialogue = 1002
	dbdMondo dbDialogue = 1003
	dbdSqlit dbDialogue = 1004
	dbdNosql dbDialogue = 1005
	dbdGaeds dbDialogue = 1006
	dbdOther dbDialogue = 1007
)

//This will be read from a text file eventually
const CURRENTDBDIALOGUE dbDialogue = dbdPg
