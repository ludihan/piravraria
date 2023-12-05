package cli

import (
	"database/sql"
	"fmt"
)

func ShowMenu(db *sql.DB) {
    fmt.Println("Bem-vindo ao terminal da piravraria.")

    fmt.Println(`Opções:
    1 - Listar entidades
    2 - (CSV) Inserir entidades
    3 - (CSV) Expotar entidades
    4 - Sair do program
    `)
}
