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

	smt, _ := db.Prepare("update usuarios set nome = ?  where id = ?")
	smt.Exec("Óxito", 1)
	smt.Exec("Uílha", 2)
}
