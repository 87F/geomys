package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = ""
	dbname   = "geomys"
)

func main() {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)

	insertStmt := `INSERT INTO "public"."Links"("url") VALUES('https://avanier.studio') RETURNING "id", "url";`
	_, e := db.Exec(insertStmt)
	CheckError(e)

	defer db.Close()
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
