package main

import (
    "bufio"
    "os"
    "strings"
	"fmt"
)

func getGuess() string {
    reader := bufio.NewReader(os.Stdin)
    fmt.Println()
    fmt.Print("Devinez une lettre: ")
    guess, _ := reader.ReadString('\n')
    fmt.Println()
    return strings.TrimSpace(guess)
}
