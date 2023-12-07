package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"piravraria/cli"

	_ "github.com/lib/pq"
)

func main() {
	connStr := "dbname=piravraria sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}

	var autenticado = cli.Autenticar(db)

	if !autenticado {
		fmt.Println("Saindo da aplicação")
		os.Exit(0)
	}

	cli.ShowMenu(db)
}
