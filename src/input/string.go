package input

import (
	"bufio"
	"fmt"
	"os"
)

func GetString(prompt string) string  {
    scanner := bufio.NewScanner(os.Stdin)
    fmt.Print(prompt)

    if scanner.Scan() {
        return scanner.Text()
    }
    return ""
}
