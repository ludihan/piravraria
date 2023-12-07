package csvlivros

import (
	"context"
	"database/sql"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"piravraria/input"
	"strconv"
)

func ImportarLivros(db *sql.DB) {
    var file = input.GetFiles()
    f, err := os.Open(file)
    if err != nil {
        log.Fatal(err)
    }

    defer f.Close()

    csvReader := csv.NewReader(f)
    data, err := csvReader.ReadAll()
    if err != nil {
        fmt.Println("erro2")
        log.Fatal(err)
    }

    ctx := context.Background()
    tx, err := db.BeginTx(ctx, nil)
    if err != nil {
        fmt.Println("erro3")
        log.Fatal(err)
    }

    for _, record := range data {
        var editora_id, _= strconv.Atoi(record[1])
        var preco_base, _ = strconv.ParseFloat(record[8],64)

        query := fmt.Sprintf(
            `INSERT INTO livro
            (isbn, editora_id, titulo, titulo_original, classificacao_etaria, descricao, edicao, idioma, preco_base)
            VALUES
            ('%s',%d,'%s','%s','%s','%s','%s','%s',%f)`,
            record[0], editora_id, record[2], record[3], record[4], record[5], record[6], record[7], preco_base,
        )
        _, err = tx.ExecContext(ctx, query)
        if err != nil {
            tx.Rollback()
            fmt.Println("\n", (err), "\n ....Rollback da transação!\n")
            return
        }

        err = tx.Commit()
        if err != nil {
            fmt.Println("erro80")
            log.Fatal(err)
        } else {
            fmt.Println("....Commit da transação\n")
        }
    }

}
