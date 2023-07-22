package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
    fmt.Print("Enter your name: ")

    scanner := bufio.NewScanner(os.Stdin)
    var name string
    if scanner.Scan() {
        name = strings.TrimSpace(scanner.Text())
        if name == "" {
            name = "World"
        }
    }
    fmt.Printf("Hello, %s!\n", name)
}

