package main

import (
	"fmt"
	"os"
)

func main() {

	printTitle()
	printMenu()

}

func printTitle() {

	title := "  _____           _ _       \n" +
		" / ____|         | | |      \n" +
		"| |  __  ___   __| | | ___  \n" +
		"| | |_ |/ _ \\ / _` | |/ _ \\ \n" +
		"| |__| | (_) | (_| | |  __/ \n" +
		" \\_____|\\___/ \\__,_|_|\\___| \n"

	fmt.Println(title)
}

// Prints the menu of game options.
func printMenu() {

	fmt.Println("Options\t\tKey")
	fmt.Println("-------\t\t---")
	fmt.Println("New Game\t n")
	fmt.Println("Rules\t\t r")
	fmt.Println("Quit\t\t q")
	fmt.Println()

	var input string
	fmt.Print("Command: ")
	fmt.Scanln(&input)

	switch input {
	case "n":
		// Start new game
	case "r":
		// Display rules, loop back to start of input
	case "q":
		fmt.Println("Thanks for playing!")
		os.Exit(0)
	default:
		fmt.Println("Invalid option.")
		// loop back to start of input
	}

}

func play() {

}

func selectWord() string {

	return ""
}
