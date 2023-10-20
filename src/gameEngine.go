package main

import (
    "math/rand"
    "time"
    "fmt"
    "strings"
)

type Game struct {
	Ligne         int
	Colonne       int
	Tableau       [][]string
	Stage         int
	Random_number int
	Mot_secret    []string
	Content       []byte
	Err           error
	Lines         []string
	Tableau_rune  []rune
	mot           string
	hasWon        bool
	hasLost       bool
	wordToGuess	  string
	guesses     []string
	attemptsLeft int
    displayedWord string
}

var g Game

func (g *Game) initializeGame() string {
    words := readWordsFromFile("words.txt")
    rand.Seed(time.Now().UnixNano())
    return words[rand.Intn(len(words))]
}

func (g *Game) Win(wordToGuess string, guesses []string) bool {
    return g.displayWord(wordToGuess, guesses) == wordToGuess
    }


func hasLost(attemptsLeft int) bool {
    return attemptsLeft <= 0
}

func (g *Game) verif() {
	fmt.Println("Jeu du Pendu")
	wordToGuess := g.initializeGame()
	guesses := []string{}
	attemptsLeft := 50

	g.displayGame(wordToGuess, guesses, attemptsLeft)
	if g.Win(wordToGuess, guesses) {
		fmt.Print("\n\n\n\n\n\n\n\n")
		fmt.Println(" __     ______  _    _  __          _______ _   _ ")
		fmt.Println(" \\ \\   / / __ \\| |  | | \\ \\        / /_   _| \\ | |")
		fmt.Println("  \\ \\_/ / |  | | |  | |  \\ \\  /\\  / /  | | |  \\| |")
		fmt.Println("   \\   /| |  | | |  | |   \\ \\/  \\/ /   | | | . ` |")
		fmt.Println("    | | | |__| | |__| |    \\  /\\  /   _| |_| |\\  |")
		fmt.Println("    |_|  \\____/ \\____/      \\/  \\/   |_____|_| \\_|")
		fmt.Println("\n\n\n")
	}
	if hasLost(attemptsLeft) {
		fmt.Println("Désolé, vous avez perdu. Le mot était:", wordToGuess)
	}

	guess := getGuess()
	guesses = append(guesses, guess)
	if !strings.Contains(wordToGuess, guess) {
		attemptsLeft--
	}
}
