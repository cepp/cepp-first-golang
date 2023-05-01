package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	db, err := sql.Open("mysql", "<USERNAME>:<PASSWORD>@/cursogo")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	tx, _ := db.Begin()

	smt, _ := tx.Prepare("insert into usuarios(id, nome) values(?, ?)")
	smt.Exec(2000, "Carlos")
	smt.Exec(2001, "Eduardo")

	_, err = smt.Exec(1, "Pedro") // chave duplicada

	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	tx.Commit()

}
