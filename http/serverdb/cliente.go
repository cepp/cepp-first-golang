package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// Usuario estrutura do usuário do sistema
type Usuario struct {
	ID   int
	Nome string
}

// UsuarioHandler respresentação dos dados do usuário
func UsuarioHandler(writer http.ResponseWriter, request *http.Request) {
	sid := strings.TrimPrefix(request.URL.Path, "/usuarios/")
	id, _ := strconv.Atoi(sid)

	switch {
	case request.Method == "GET" && id > 0:
		usuarioPorID(writer, request, id)
	case request.Method == "GET":
		usuarioTodos(writer, request)
	default:
		writer.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(writer, "Desculpa... :(")
	}
}

func usuarioPorID(writer http.ResponseWriter, request *http.Request, id int) {
	db, err := sql.Open("mysql", "<USERNAME>:<PASSWORD>@/cursogo")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var usuario Usuario
	db.QueryRow("select id, nome from usuarios where id > ?", id).Scan(&usuario.ID, &usuario.Nome)

	responseJson, _ := json.Marshal(usuario)

	writer.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(writer, string(responseJson))
}

func usuarioTodos(writer http.ResponseWriter, request *http.Request) {
	db, err := sql.Open("mysql", "<USERNAME>:<PASSWORD>@/cursogo")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, _ := db.Query("select id, nome from usuarios")
	defer rows.Close()

	var usuarios []Usuario
	for rows.Next() {
		var usuario Usuario
		rows.Scan(&usuario.ID, &usuario.Nome)
		usuarios = append(usuarios, usuario)
	}
	responseJson, _ := json.Marshal(usuarios)

	writer.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(writer, string(responseJson))
}
