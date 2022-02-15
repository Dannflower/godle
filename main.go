package main

import (
	"bufio"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/fatih/color"
)

const (
	// The letter is not in the word.
	notInWord int = iota
	// The letter is in the word and in the correct position.
	correctPosition
	// The character is in the word but in the wrong position.
	wrongPosition
)

const maxGuesses int = 6

var scanner *bufio.Scanner

func init() {
	rand.Seed(int64(time.Now().Nanosecond()))
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
	fmt.Printf("You get %v guesses to get the right word.\n", maxGuesses)
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
	fmt.Printf("Guesses: %v/%v\n", len(guesses), maxGuesses)
	fmt.Println("Hit enter to return to the menu.")
	scanner.Scan()
}

// Start the core game loop.
func play() {

	answer := selectWord()
	var guesses []string
	var results [][]int
	usedLetters := make(map[rune]int)

	fmt.Println("Guess the word!")

	for len(guesses) < maxGuesses {

		fmt.Print("Guess: ")
		scanner.Scan()

		guess := scanner.Text()
		result, err := compareRunes(convertToRunes(guess), convertToRunes(answer), usedLetters)

		if err != nil {

			fmt.Printf("Guesses must be %v characters long.\n", len(answer))

		} else {

			guesses = append(guesses, guess)
			results = append(results, result)
			printGuessResult(guesses, results)
			printAvailableLetters(usedLetters)

			// Player has won!
			if guess == answer {

				handleWin(guesses)
				return
			}
		}
	}

	fmt.Printf("Nice try! The word was '%s.'\n", answer)
	fmt.Println("Hit enter to return to the menu.")
	scanner.Scan()
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

	case notInWord:
		return color.HiBlackString(str)

	case wrongPosition:
		return color.YellowString(str)

	case correctPosition:
		return color.GreenString(str)

	default:
		return str
	}
}

// Prints the results of the last guess and all previous guesses
// with runes color coded depending on whether they are in the word,
// not in the word, or in the word but the wrong location.
func printGuessResult(guesses []string, results [][]int) {

	for i, guess := range guesses {

		capGuess := strings.ToUpper(guess)
		colorResult := ""

		for j, r := range capGuess {

			colorResult += addHintColor(string(r), results[i][j])
		}

		fmt.Println(colorResult)
	}
}

// Converts the given string into a slice of runes.
// Each rune is the upper case.
func convertToRunes(word string) []rune {

	var runes []rune

	for _, r := range strings.ToUpper(word) {

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
func compareRunes(guess []rune, answer []rune, usedLetters map[rune]int) ([]int, error) {

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
			usedLetters[r] = correctPosition
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

			// Don't overwrite correct position status on letters
			if usedLetters[r] != correctPosition {

				usedLetters[r] = wrongPosition
			}
		}

		// Mark any letters that aren't in the word
		if usedLetters[r] != wrongPosition && usedLetters[r] != correctPosition {

			usedLetters[r] = notInWord
		}
	}

	return result, nil
}

// Selects a random word from the list of answer words.
func selectWord() string {

	return Words[rand.Intn(len(Words))]
}
