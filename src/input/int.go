package input

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func GetInt(prompt string) int  {
    scanner := bufio.NewScanner(os.Stdin)
    fmt.Print(prompt)

    var buffer string
    for {

        if scanner.Scan() {
            buffer = scanner.Text()
        }
        if i, err := strconv.Atoi(buffer); err != nil {
            fmt.Println(err)
            continue
        } else {
            return i
        }
    }
}
