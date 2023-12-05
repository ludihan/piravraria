package cli

import (
	"database/sql"
	"fmt"
	"os"
)

func ShowMenu(db *sql.DB) {
    fmt.Println("Bem-vindo ao terminal da piravraria.")

    for {
        fmt.Println(`Opções:
        1 - Listar entidades
        2 - (CSV) Inserir entidades
        3 - (CSV) Expotar entidades
        4 - Sair do program
        `)
        var input int

        fmt.Scanln(&input)

        switch input {
        case 1:
            listarClientesEEndereco(db)
        case 2:
        case 3:
        case  4:
            os.Exit(0)
        default:
            fmt.Println("Não é uma opção válida.")
        }

    }

}
