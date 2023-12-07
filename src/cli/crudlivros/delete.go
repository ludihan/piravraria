package crudlivros

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"piravraria/input"
)

func DeleteLivros(db *sql.DB) {
    var isbn string
    isbn = input.GetString("ISBN do livro que vai ser deletado:\n$ ")
    ctx := context.Background()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
    query := fmt.Sprintf(
        `DELETE FROM livro
        WHERE
        isbn = '%s'`,
        isbn,
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
