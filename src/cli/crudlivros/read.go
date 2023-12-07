package crudlivros

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"piravraria/input"

	"github.com/olekukonko/tablewriter"
)

func ReadLivros(db *sql.DB) {
    var isbn string
    isbn = input.GetString("ISBN do livro que vai ser listado (apenas aperte enter para mostrar todos):\n$ ")

    var query string
    if isbn == "" {
        query = `SELECT * FROM livro`
    } else {
        query = fmt.Sprintf(
            `SELECT * FROM livro
            WHERE
            isbn = '%s'`,
            isbn,
        )
    }

    rows, err := db.Query(query)
    defer rows.Close()
    if err != nil {
        log.Fatal(err)
    }
    table := tablewriter.NewWriter(os.Stdout)
    table.SetHeader([]string{"isbn", "editora_id", "titulo", "titulo_original", "classificacao_etaria", "descricao", "edicao", "idioma", "preco_base"})

    var titulo, classificacao_etaria string
    var preco_base float64
    var editora_id sql.NullInt64
    var titulo_original, descricao, edicao, idioma sql.NullString
    for rows.Next() {
        err = rows.Scan(&isbn, &editora_id, &titulo, &titulo_original, &classificacao_etaria, &descricao, &edicao, &idioma, &preco_base)
        if err != nil {
            fmt.Println(err)
        }
        table.Append([]string {isbn, nullInt(editora_id), titulo, nullString(titulo_original), classificacao_etaria, nullString(descricao), nullString(edicao),
        nullString(idioma), fmt.Sprint(preco_base)})
    }
    table.Render()
}

func nullString(value sql.NullString) string {
    if value.Valid {
        return value.String
    } else {
        return ""
    }
}

func nullInt(value sql.NullInt64) string {
    if value.Valid {
        return fmt.Sprintf("%v",value.Int64)
    } else {
        return ""
    }
}
