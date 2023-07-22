package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func main() {
    running := true
    scanner := bufio.NewScanner(os.Stdin)

    for running {
        fmt.Println("=== Guessing Game ===")

        number := int64(rand.Intn(101))
        success := false
        fmt.Println("Number generated between 0 to 100")
        for i := 0; i < 10; i++ {
            fmt.Printf("Remaining guesses: %d\n", 10 - i)
            fmt.Print("Enter your guess: ")
            scanner.Scan()
            guess, err := strconv.ParseInt(scanner.Text(), 10, 64)
            if err != nil {
                fmt.Println("Invalid guess. Guess again.")
                continue
            }
            if guess > number {
                fmt.Println("Too high!")
            } else if guess < number {
                fmt.Println("Too low!")
            } else {
                fmt.Println("You guessed right!")
                success = true
                break
            }
        }

        if success {
            fmt.Println("Congratulations!")
        } else {
            fmt.Println("You ran out of guesses...")
        }
        fmt.Print("Try again? [y/N]: ")
        scanner.Scan()
        retry := strings.TrimSpace(scanner.Text())
        if retry == "y" || retry == "Y" {
            continue
        } else {
            running = false
        }
    }

    fmt.Println("Goodbye!")
}
