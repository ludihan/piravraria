package csvlivros

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"
)

func ExportarLivros(db *sql.DB) {
    f, err := os.Create("livros " + time.Now().Format("01-02-2006 15:04:05 Monday") + ".csv")
    if err != nil {
        log.Fatal(err)
    }
    defer f.Close()
    fmt.Println("criando arquivo com nome de " + f.Name())

    /*
    _, err2 := f.WriteString("isbn,editora_id,titulo,titulo_original,classificacao_etaria,descricao,edicao,idioma,preco_base\n")
    if err2 != nil {
        log.Fatal(err2)
    }
    */

    rows, err := db.Query("SELECT * FROM livro")
    defer rows.Close()
    if err != nil {
        log.Fatal(err)
    }
    var isbn, titulo, classificacao_etaria string
    var preco_base float64
    var editora_id sql.NullInt64
    var titulo_original, descricao, edicao, idioma sql.NullString
    for rows.Next() {
        err = rows.Scan(&isbn, &editora_id, &titulo, &titulo_original, &classificacao_etaria, &descricao, &edicao, &idioma, &preco_base)
        if err != nil {
            fmt.Println(err)
        }
        _, err2 := f.WriteString(isbn + "," + nullInt(editora_id) + "," + titulo + "," + nullString(titulo_original) + "," + classificacao_etaria + "," +
        nullString(descricao) + "," + nullString(edicao) + "," + nullString(idioma) + "," + fmt.Sprint(preco_base) + "\n")
        if err2 != nil {
            log.Fatal(err2)
        }
    }
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
