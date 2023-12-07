package input

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func GetFiles() string {
    files, err := os.ReadDir(".")
    if err != nil {
        log.Fatal(err)
    }
    var foundFiles = []string{}

    var i int
    for _, file := range files {
        if !file.IsDir() {
            fmt.Println(i, "=", file.Name())
            foundFiles = append(foundFiles, file.Name())
            i++
        }
    }


    scanner := bufio.NewScanner(os.Stdin)

    var buffer string
    for {
        fmt.Print("$ ")

        if scanner.Scan() {
            buffer = scanner.Text()
        }
        if i, err := strconv.Atoi(buffer); err != nil {
            fmt.Println(err)
            continue
        } else {
            return foundFiles[i]
        }
    }
}
