package cli

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
)

type Admin struct {
    nome string
    senha string
}

func Autenticar(db *sql.DB) bool {
    fmt.Print("Nome:  ")
    var nome, senha string

    if _, err := fmt.Scanln(&nome); err != nil {
        log.Fatal(err)
    }

    fmt.Print("Senha: ")
    if _, err := fmt.Scanln(&senha); err != nil {
        log.Fatal(err)
    }

    nome = strings.TrimSpace(nome)
    senha = strings.TrimSpace(senha)

    rows, err := db.Query("SELECT nome, senha FROM admin WHERE nome = $1 AND senha = $2", nome, senha)
    if err != nil {
        log.Fatal(err)
    }

    if rows.Next() {
        fmt.Printf("O usuário com nome %v e senha %v foi autenticado\n", nome, senha)
        return true
    } else {
        fmt.Printf("O usuário com nome %v e senha %v não possui permissão ao sistema\n", nome, senha)
        return false
    }
}
