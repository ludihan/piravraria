package cli

import (
    "database/sql"
    "log"
    "os"

    "github.com/olekukonko/tablewriter"
)

func selectClienteEndereco(db *sql.DB) {
    rows, err := db.Query("SELECT nome, email, cep, uf, localidade, logradouro, numero, complemento FROM cliente INNER JOIN endereco ON cliente.endereco_id = endereco.id;")
    defer rows.Close()
    if err != nil {
        log.Fatal(err)
    }
    table := tablewriter.NewWriter(os.Stdout)
    table.SetHeader([]string{"nome", "email", "cep", "uf", "localidade", "logradouro", "numero", "complemento"})

    var nome, email, cep, uf, localidade, logradouro, numero, complemento string
    for rows.Next() {
        rows.Scan(&nome, &email, &cep, &uf, &localidade, &logradouro, &numero, &complemento)
        table.Append([]string {nome, email, cep, uf, localidade, logradouro, numero, complemento})
    }
    table.Render()
}
