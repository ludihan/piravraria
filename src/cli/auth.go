package cli

import (
    "piravraria/input"
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
    var nome, senha string

    nome = input.GetString("Nome: ")
    senha = input.GetString("Senha: ")


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
