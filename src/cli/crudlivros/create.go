package crudlivros

import (
    "piravraria/input"
	"context"
	"database/sql"
	"fmt"
	"log"
)

func CreateLivros(db *sql.DB) {
    var isbn, editora_id, titulo, titulo_original, classificacao_etaria, descricao, edicao, idioma, preco_base string

    isbn = input.GetString("isbn:\n$ ")
    fmt.Println("\n(Esse valor precisa ser um int)")
    editora_id = input.GetString("editora_id:\n$ ")
    titulo = input.GetString("titulo:\n$ ")
    titulo_original = input.GetString("titulo_original:\n$ ")
    fmt.Println("\n(Esse valor precisa ser l', '10', '12', '14', '16' ou '18')")
    classificacao_etaria = input.GetString("classificacao_etaria:\n$ ")
    descricao = input.GetString("descricao:\n$ ")
    edicao = input.GetString("edicao:\n$ ")
    idioma = input.GetString("idioma:\n$ ")
    fmt.Println("\n(Esse valor precisa ser um int)")
    preco_base = input.GetString("preco_base:\n$ ")

    ctx := context.Background()
    tx, err := db.BeginTx(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }

    query := fmt.Sprintf(
        `INSERT INTO livro
        (isbn, editora_id, titulo, titulo_original, classificacao_etaria, descricao, edicao, idioma, preco_base)
        VALUES
        ('%s','%s','%s','%s','%s','%s','%s','%s','%s')`,
        isbn, editora_id, titulo, titulo_original, classificacao_etaria, descricao, edicao, idioma, preco_base,
    )
    _, err = tx.ExecContext(ctx, query)
    if err != nil {
        tx.Rollback()
        fmt.Println("\n", (err), "\n ....Rollback da transação!\n")
        return
    }

    err = tx.Commit()
    if err != nil {
        log.Fatal(err)
    } else {
        fmt.Println("....Commit da transação\n")
    }
}
