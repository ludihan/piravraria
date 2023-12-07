package cli;

import (
    "database/sql"
    "fmt"
    "os"
)

func ShowMenu(db *sql.DB) {
    fmt.Println("Bem-vindo ao terminal da piravraria.")

    var input string

    for {
        fmt.Print(
            "Opções:\n",
            "1 - (SELECT)     Listar entidades\n",
            "2 - (INSERT)     Inserir entidades\n",
            "3 - (CSV IMPORT) Inserir entidades\n",
            "4 - (CSV EXPORT) Expotar entidades\n",
            "5 - Sair do program\n",
            "$ ",
        )

        if _, err := fmt.Scanln(&input); err != nil {
            fmt.Println("\n", err)
            continue
        }

        switch input {
        case "1":
            menuSelect(db)
        case "2":
            menuInsert(db)
        case "3":
            menuCSVImport(db)
        case "4":
            menuCSVExport(db)
        case "5":
            os.Exit(0)
        default:
            fmt.Println("Não é uma opção válida.")
        }
    }
}

func menuSelect(db *sql.DB) {

    var input string

    for {
        fmt.Print(
            "SELECT:\n",
            "1 - Voltar\n",
            "2 - Cliente e endereço\n",
            "3 - Livro e autor\n",
            "4 - Livro\n",
            "$ ",
        )

        if _, err := fmt.Scanln(&input); err != nil {
            fmt.Println("\n", err)
            continue
        }

        switch input {
        case "1":
            return
        case "2":
            selectClienteEndereco(db)
        case "3":
            selectLivroAutor(db)
        case "4":
            selectLivro(db)
        default:
            fmt.Println("Não é uma opção válida.")
        }
    }
}

func menuInsert(db *sql.DB) {
    var input string

    for {
        fmt.Print(
            "INSERT:\n",
            "1 - Voltar\n",
            "2 - Livro\n",
            "$ ",
        )

        if _, err := fmt.Scanln(&input); err != nil {
            fmt.Println("\n", err)
            continue
        }

        switch input {
        case "1":
            return
        case "2":
            insertLivro(db)
        default:
            fmt.Println("Não é uma opção válida.")
        }
    }

}

func menuCSVImport(db *sql.DB) {
    var input string

    for {
        fmt.Print(
            "SELECT:\n",
            "1 - Cliente e endereço\n",
            "2 - Livro e autor\n",
            "3 - Livro e categoria\n",
            "4 - Voltar\n",
            "$ ",
        )

        if _, err := fmt.Scanln(&input); err != nil {
            fmt.Println("\n", err)
            continue
        }

        switch input {
        case "1":
            selectClienteEndereco(db)
        case "2":
            menuInsert(db)
        case "3":
            menuCSVImport(db)
        case "4":
            return
        default:
            fmt.Println("Não é uma opção válida.")
        }
    }

}

func menuCSVExport(db *sql.DB) {
    var input string

    for {
        fmt.Print(
            "SELECT:\n",
            "1 - Cliente e endereço\n",
            "2 - Livro e autor\n",
            "3 - Livro e categoria\n",
            "4 - Voltar\n",
            "$ ",
        )

        if _, err := fmt.Scanln(&input); err != nil {
            fmt.Println("\n", err)
            continue
        }

        switch input {
        case "1":
            selectClienteEndereco(db)
        case "2":
            menuInsert(db)
        case "3":
            menuCSVImport(db)
        case "4":
            return
        default:
            fmt.Println("Não é uma opção válida.")
        }
    }

}
