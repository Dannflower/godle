package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"time"
)

const (
	// The letter is not in the word.
	notInWord int = iota
	// The letter is in the word and in the correct position.
	correctPosition
	// The character is in the word but in the wrong position.
	wrongPosition
)

func init() {
	rand.Seed(int64(time.Now().Nanosecond()))
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

		var input string

		fmt.Print("Command: ")
		fmt.Scanln(&input)

		switch input {
		case "p":
			// Start new game
			play()
			printMenu()
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
}

func handleWin() {

	fmt.Println("You got it! Hit enter to return to the menu.")
	fmt.Scanln()
}

// Start the core game loop.
func play() {

	answer := selectWord()

	fmt.Println("Answer: " + answer)

	for {

		guess := ""

		fmt.Print("Guess: ")
		fmt.Scanln(&guess)

		result, err := compareRunes(convertToRunes(guess), convertToRunes(answer))

		if err != nil {

			fmt.Printf("Guesses must be at least %v characters long.\n", len(answer))

		} else {

			fmt.Println(result)

			// Player has won!
			if guess == answer {

				handleWin()
				return
			}
		}
	}
}

// Converts the given string into a slice of runes.
func convertToRunes(word string) []rune {

	var runes []rune

	for _, r := range word {

		runes = append(runes, r)
	}

	return runes
}

// Returns a slice containing the indices in runes in which
// r is found. If r is not in runes, nil is returned.
func getRuneIndices(runes []rune, r rune) []int {

	var indices []int

	for i, ru := range runes {

		if ru == r {
			indices = append(indices, i)
		}
	}

	return indices
}

// Compares the guess runes to the answer runes and returns
// the result as a slice of equal length.
//
// The result slice will be comprised of the enums describing
// whether each rune is in the answer, in the correct position,
// or not in the answer at all.
//
// If the guess and answer slices are of different lengths, an error is returned.
func compareRunes(guess []rune, answer []rune) ([]int, error) {

	if len(guess) != len(answer) {

		return nil, errors.New("length of guess and answer must match")
	}

	result := make([]int, len(answer))
	occurences := make(map[rune]int)

	// Check for correct letters first
	for i, r := range guess {

		if answer[i] == r {

			occurences[r] = occurences[r] + 1
			result[i] = correctPosition
		}
	}

	// Check for characters in the wrong position
	for i, r := range guess {

		// Skip runes already checked
		if result[i] == correctPosition {

			continue
		}

		occurences[r] = occurences[r] + 1
		answerIndices := getRuneIndices(answer, r)

		// Ignore extra occurences of a rune
		if occurences[r] <= len(answerIndices) {

			result[i] = wrongPosition
		}
	}

	return result, nil
}

// Selects a random word from the list of answer words.
func selectWord() string {

	return Words[rand.Intn(len(Words))]
}
