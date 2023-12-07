package crudlivros

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"piravraria/input"
)

func UpdateLivros(db *sql.DB) {
    var isbn string
    isbn = input.GetString("ISBN do livro que vai ser modificado:\n$ ")

    var prompt = "Qual valor você gostaria de mudar?\n" +
    "1 - editora_id\n" +
    "2 - titulo\n" +
    "3 - titulo_original\n" +
    "4 - classificacao_etaria\n" +
    "5 - descricao\n" +
    "6 - edicao\n" +
    "7 - idioma\n" +
    "8 - preco_base\n$ "

    var colunas = [...]string{"editora_id", "titulo", "titulo_original", "classificacao_etaria", "descricao", "edicao", "idioma", "preco_base"}

    var inputOpcao string
    inputOpcao = input.GetString(prompt)

    var colunaAlvo string
    switch inputOpcao {
    case "1":
        colunaAlvo = colunas[0]
    case "2":
        colunaAlvo = colunas[1]
    case "3":
        colunaAlvo = colunas[2]
    case "4":
        colunaAlvo = colunas[3]
    case "5":
        colunaAlvo = colunas[4]
    case "6":
        colunaAlvo = colunas[5]
    case "7":
        colunaAlvo = colunas[6]
    case "8":
        colunaAlvo = colunas[7]
    }

    var valor string
    valor = input.GetString("Qual valor?\n$ ")

    ctx := context.Background()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
    query := fmt.Sprintf(
        `UPDATE livro
        SET %s = '%s'
        WHERE isbn = '%s'`,
        colunaAlvo, valor, isbn,
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
