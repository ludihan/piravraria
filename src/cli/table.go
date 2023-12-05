package cli

import (
	"database/sql"
	"log"
	"os"

	"github.com/olekukonko/tablewriter"
)

func listarClientesEEndereco(db *sql.DB) {
    rows, err := db.Query("SELECT nome, email, cep, uf, localidade, logradouro, numero, complemento FROM cliente INNER JOIN endereco ON cliente.endereco_id = endereco.id;", nome, senha)
    defer rows.Close()
    if err != nil {
        log.Fatal(err)
    }
    table := tablewriter.NewWriter(os.Stdout)
    table.SetHeader([]string{"nome", "email", "cep", "uf", "localidade", "logradouro", "numero", "complemento"})

    var nome, email, cep, uf, localidade, logradouro, numero, complemento string
    for rows.Next() {
        queryTable := [][]string {

        }
        rows.Scan(&nome, &email, &cep, &uf, &localidade, &logradouro, &numero, &complemento)
        table.Append()
    }
}


func listarLivrosEAutor(db *sql.DB) {

}

func listarLivrosECategoria(db *sql.DB) {

}
