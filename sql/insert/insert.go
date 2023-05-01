package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "<USERNAME>:<PASSWORD>@/cursogo")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	smt, _ := db.Prepare("insert into usuarios(nome) values(?)")
	smt.Exec("Maria")
	smt.Exec("João")

	res, _ := smt.Exec("Pedro")

	id, _ := res.LastInsertId()
	fmt.Println(id)

	linhas, _ := res.RowsAffected()
	fmt.Println(linhas)
}
