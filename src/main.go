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
    if err != nil {
		log.Fatal(err)
	}


    if !cli.Autenticar(db) {
        fmt.Println("Desligando o sistema")
        os.Exit(0)
    }

    fmt.Println("Bem-vindo ao terminal da piravraria.")

}
