package cli

import (
	"database/sql"
	"fmt"
	"os"
	"piravraria/cli/crudlivros"
	"piravraria/csvlivros"
	"piravraria/input"
)

func ShowMenu(db *sql.DB) {
    fmt.Println("Bem-vindo ao terminal da piravraria.")
    var bufferInput string
    var prompt string = "Opções:\n" +
    "\t1 - CRUD para livros              (CRUD)\n" +
    "\t2 - Mostrar clientes com endereço (SELECT COM JOIN)\n" +
    "\t3 - Sair do programa              (EXIT)\n" +
    "$ "

    for {
        bufferInput = input.GetString(prompt)

        switch bufferInput {
        case "1":
            showMenuCRUD(db)
        case "2":
            selectClienteEndereco(db)
        case "3":
            fmt.Println("Vá com Deus")
            os.Exit(0)
        default:
            fmt.Println("Não é uma opção válida.")
        }
    }
}

func showMenuCRUD(db *sql.DB) {
    var bufferInput string
    var prompt string =
    "Opcões:\n" +
    "\t1 - Inserir livros            (CREATE)\n" +
    "\t2 - Mostrar livros            (READ)\n" +
    "\t3 - Mudar livros              (UPDATE)\n" +
    "\t4 - Deletar livros            (DELETE)\n" +
    "\t5 - Importar CSV              (IMPORT CSV)\n" +
    "\t6 - Exportar CSV              (EXPORT CSV)\n" +
    "\t7 - Voltar pro menu principal (EXIT)\n" +
    "$ "

    for {
        bufferInput = input.GetString(prompt)

        switch bufferInput {
        case "1":
            crudlivros.CreateLivros(db)
        case "2":
            crudlivros.ReadLivros(db)
        case "3":
            crudlivros.UpdateLivros(db)
        case "4":
            crudlivros.DeleteLivros(db)
        case "5":
            csvlivros.ImportarLivros(db)
        case "6":
            csvlivros.ExportarLivros(db)
        case "7":
            return
        default:
            fmt.Println("Não é uma opção válida.")
        }
    }
}
