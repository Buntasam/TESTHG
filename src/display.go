package main

import (
	"fmt"
	"strings"
)

func (g *Game) displayGame(wordToGuess string, guesses []string, attemptsLeft int) {
	fmt.Println("\nMot à deviner:", g.displayWord(wordToGuess, guesses))
	fmt.Println()
	fmt.Println("Tentatives restantes:", attemptsLeft)
	fmt.Println()
	fmt.Println("Lettres déjà utilisées:", strings.Join(guesses, " "))
}

func (g *Game) displayWord(wordToGuess string, guesses []string) string {
	displayedWord := ""
	for _, letter := range wordToGuess {
		if strings.Contains(strings.Join(guesses, ""), string(letter)) {
			displayedWord += string(letter)
		} else {
			displayedWord += "_"
		}
	}
	return displayedWord
}

func (g *Game) Display() {

	switch g.Stage {
	case 0: //vide
		for i := 0; i < len(g.Tableau); i++ {
			fmt.Println(g.Tableau[i])
		}
		g.verif()
		g.Win(g.wordToGuess, g.guesses)
		g.Display()

	case 1: //sol
		fmt.Print("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n")
		for i := 0; i < 9; i++ {
			g.Tableau[g.Ligne-1][i] = "-"
		}

		for i := 0; i < len(g.Tableau); i++ {
			fmt.Println(g.Tableau[i])
		}

		g.verif()
		g.Win(g.wordToGuess, g.guesses)
		g.Display()

	case 2: //poteau
		fmt.Print("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n")
		for i := 0; i < 7; i++ {
			g.Tableau[i][4] = "|"
		}
		for i := 0; i < len(g.Tableau); i++ {
			fmt.Println(g.Tableau[i])
		}
		g.verif()
		g.Win(g.wordToGuess, g.guesses)
		g.Display()

	case 3: //barre du haut
		fmt.Print("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n")
		for i := 4; i < 12; i++ {
			g.Tableau[0][i] = "-"
		}
		for i := 0; i < len(g.Tableau); i++ {
			fmt.Println(g.Tableau[i])
		}

		g.verif()
		g.Win(g.wordToGuess, g.guesses)
		g.Display()

	case 4: //soutien
		fmt.Print("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n")
		g.Tableau[1][6] = "/"
		g.Tableau[2][5] = "/"

		for i := 0; i < len(g.Tableau); i++ {
			fmt.Println(g.Tableau[i])
		}

		g.verif()
		g.Win(g.wordToGuess, g.guesses)
		g.Display()

	case 5: //corde
		fmt.Print("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n")
		g.Tableau[1][11] = "|"

		for i := 0; i < len(g.Tableau); i++ {
			fmt.Println(g.Tableau[i])
		}

		g.verif()
		g.Win(g.wordToGuess, g.guesses)
		g.Display()

	case 6: //tete
		fmt.Print("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n")
		g.Tableau[2][11] = "O"

		for i := 0; i < len(g.Tableau); i++ {
			fmt.Println(g.Tableau[i])
		}

		g.verif()
		g.Win(g.wordToGuess, g.guesses)
		g.Display()

	case 7: // corps
		fmt.Print("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n")
		g.Tableau[3][11] = "|"
		g.Tableau[4][11] = "|"

		for i := 0; i < len(g.Tableau); i++ {
			fmt.Println(g.Tableau[i])
		}

		g.verif()
		g.Win(g.wordToGuess, g.guesses)
		g.Display()


	case 8: // bras droit
		fmt.Print("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n")
		g.Tableau[3][12] = `\`

		for i := 0; i < len(g.Tableau); i++ {
			fmt.Println(g.Tableau[i])
		}

		g.verif()
		g.Win(g.wordToGuess, g.guesses)
		g.Display()
		

	case 9: // bras gauche
		fmt.Print("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n")
		g.Tableau[3][10] = "/"

		for i := 0; i < len(g.Tableau); i++ {
			fmt.Println(g.Tableau[i])
		}

		g.verif()
		g.Win(g.wordToGuess, g.guesses)
		g.Display()


	case 10: // jambe gauche
		fmt.Print("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n")
		g.Tableau[5][10] = "/"

		for i := 0; i < len(g.Tableau); i++ {
			fmt.Println(g.Tableau[i])
		}

		g.verif()
		g.Win(g.wordToGuess, g.guesses)
		g.Display()
		

	case 11: // jambe droite
		fmt.Print("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n")
		g.Tableau[5][12] = `\`

		for i := 0; i < len(g.Tableau); i++ {
			fmt.Println(g.Tableau[i])
		}

		g.verif()
		g.Win(g.wordToGuess, g.guesses)
		g.Display()
		

	case 12: // game over
		fmt.Print("\n\n")
		g.Tableau[5][12] = `\`
		fmt.Println("  ________                        ________                     ")
		fmt.Println(" /  _____/ ____    _____   ____  ", `\`+"_____ ", `\`+"___  __ ___________ ")
		fmt.Println("/  ", `\`, " ___", `\`+"__", `\`, " /    ", `\`+"_/ __", `\`, "  /   |  ", `\`+" ", `\`+"/ // __", `\`+"_  __", `\`)
		fmt.Println(`\`, "  ", `\`+"_"+`\ `, `\`+"/ __", `\`+"|  Y Y ", `\`, " ___/  /    |   ", `\`, "  /", `\`, "___/|  | "+`\`+"/")
		fmt.Println(` \` + "______  (____  /__|_|  /" + `\` + "___  > " + `\` + "_______  /" + `\` + "_/   " + `\` + "___ >__|   ")
		fmt.Println("       ", `\`+"/    ", `\`+"/      "+`\`+"/    ", `\`+"/         ", `\`+"/           "+`\`+"/       ")
		fmt.Println()
		fmt.Println()
		fmt.Printf("Le mot était : %s", g.mot)
		fmt.Println()
		fmt.Println()
	}
}
