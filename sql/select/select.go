package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type usuario struct {
	id   int
	nome string
}

func main() {
	db, err := sql.Open("mysql", "<USERNAME>:<PASSWORD>@/cursogo")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, _ := db.Query("select id, nome from usuarios where id > ?", 5)
	defer rows.Close()

	for rows.Next() {
		var usuario usuario
		rows.Scan(&usuario.id, &usuario.nome)
		fmt.Println(usuario)
	}
}
