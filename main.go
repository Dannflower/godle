package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Dannflower/godle/logic"

	"github.com/fatih/color"
)

var scanner *bufio.Scanner

func init() {
	scanner = bufio.NewScanner(os.Stdin)
}

func main() {

	printTitle()
	printMenu()
	handleMenuInput()
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

func printMenu() {

	fmt.Println("Options\t\tKey")
	fmt.Println("-------\t\t---")
	fmt.Println("Play\t\t p")
	fmt.Println("Rules\t\t r")
	fmt.Println("Quit\t\t q")
	fmt.Println()
}

// Handles menu option selections
func handleMenuInput() {

	for {

		fmt.Print("Command: ")
		scanner.Scan()

		switch scanner.Text() {
		case "p":
			// Start new game
			play()
			printMenu()
		case "r":
			// Display rules, loop back to start of input
			printRules()
			printMenu()
		case "q":
			fmt.Println("Thanks for playing!")
			os.Exit(0)
		default:
			fmt.Println("Invalid option.")
			// loop back to start of input
		}
	}
}

func printRules() {
	fmt.Println("Attempt to guess a randomly selected 5-letter word.")
	fmt.Printf("You get %v guesses to get the right word.\n", logic.MaxGuesses)
	fmt.Println("After guessing your guess will be displayed with color coding indicating the following:")
	color.HiBlack("Gray - The letter is not in the word.")
	color.Yellow("Yellow - The letter is in the word but is in the wrong position.")
	color.Green("Green - The letter is in the word and in the right position.")
	fmt.Println("If all guesses are exhausted, the answer will be revealed. Good luck word nerd!")
	fmt.Println("Hit enter to return to the menu.")
	scanner.Scan()
}

func handleWin(guesses []string) {

	fmt.Println("You got it!")
	fmt.Printf("Guesses: %v/%v\n", len(guesses), logic.MaxGuesses)
	fmt.Println("Hit enter to return to the menu.")
	scanner.Scan()
}

// Start the core game loop.
func play() {

	fmt.Println("Guess the word!")

	for len(logic.Guesses) < logic.MaxGuesses {

		fmt.Print("Guess: ")
		scanner.Scan()

		guess := scanner.Text()
		err := logic.MakeGuess(guess)

		if err != nil {

			fmt.Printf("Guesses must be %v characters long.\n", len(logic.Answer))

		} else {

			printGuessResult()
			printAvailableLetters(logic.UsedLetters)

			// Player has won!
			if logic.HasWon(guess) {

				handleWin(logic.Guesses)
				return
			}
		}
	}

	fmt.Printf("Nice try! The word was '%s.'\n", logic.Answer)
	fmt.Println("Hit enter to return to the menu.")
	scanner.Scan()
}

// Prints the results of the last guess and all previous guesses
// with runes color coded depending on whether they are in the word,
// not in the word, or in the word but the wrong location.
func printGuessResult() {

	for i, guess := range logic.Guesses {

		capGuess := strings.ToUpper(guess)
		colorResult := ""

		for j, r := range capGuess {

			colorResult += addHintColor(string(r), logic.Results[i][j])
		}

		fmt.Println(colorResult)
	}
}

// Prints out the complete list of letters with any
// used in previous guesses printed in gray.
func printAvailableLetters(usedLetters map[rune]int) {

	letters := []rune{
		'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M',
		'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}

	line := ""

	for i, letter := range letters {

		colorLetter := string(letter)

		// Add hint color if the letter has been used
		if _, ok := usedLetters[letter]; ok {
			colorLetter = addHintColor(colorLetter, usedLetters[letter])
		}

		line += colorLetter + " "

		if i == 12 || i == 25 {

			fmt.Println(line)
			line = ""
		}
	}
}

// Returns the ANSI coded version of the string
// colored appropriately for the given hint.
func addHintColor(str string, hint int) string {

	switch hint {

	case logic.NotInWord:
		return color.HiBlackString(str)

	case logic.WrongPosition:
		return color.YellowString(str)

	case logic.CorrectPosition:
		return color.GreenString(str)

	default:
		return str
	}
}
