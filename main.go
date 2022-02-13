package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strings"
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
	requestMenuInput()
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

// Prints the menu of game options.
func requestMenuInput() {

	for {

		var input string

		fmt.Print("Command: ")
		fmt.Scanln(&input)

		switch input {
		case "p":
			// Start new game
			play()
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

// Start the core game loop.
func play() {

	// Test
	// maxGuesses := 6
	answer := selectWord()

	fmt.Println("Answer: " + answer)

	for {

		guess := ""

		fmt.Print("Guess: ")
		fmt.Scanln(&guess)

		result, err := compareWords(guess, answer)

		if err != nil {

			fmt.Printf("Guesses must be at least %v characters long.\n", len(answer))

		} else {

			fmt.Println(result)
		}
	}
}

// Compares the guess string to the answer string and returns
// the result as a slice of equal length.
//
// The result slice will be comprised of the enums describing
// whether each character is in the word, in the correct position,
// or not in the word.
//
// If the guess and answer strings are of different lengths, an error is returned.
func compareWords(guess string, answer string) ([]int, error) {

	if len(guess) != len(answer) {

		return nil, errors.New("length of guess and answer must match")
	}

	result := make([]int, len(answer))

	for i, c := range guess {

		answerRuneIndex := strings.IndexRune(answer, c)

		switch {

		case i == answerRuneIndex:
			result[i] = correctPosition

		case answerRuneIndex >= 0:
			// TODO: Handle double letters
			result[i] = wrongPosition

		case answerRuneIndex == -1:
			result[i] = notInWord
		}
	}

	return result, nil
}

func selectWord() string {

	return Words[rand.Intn(len(Words))]
}
