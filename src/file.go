package main

import (
    "os"
    "strings"
	"fmt"
)

func readWordsFromFile(liste string) []string {
    content, err := os.ReadFile("/home/zeus/Bureau/ZTESTGP/src/mot.txt")
    if err != nil {
        fmt.Println("Erreur lors de la lecture du fichier:", err)
        return nil
    }
    return strings.Fields(string(content))
}
