package logic

import (
	"errors"
	"math/rand"
	"strings"
	"time"
)

const MaxGuesses int = 6

const (
	// The letter is not in the word.
	NotInWord int = iota
	// The letter is in the word and in the correct position.
	CorrectPosition
	// The character is in the word but in the wrong position.
	WrongPosition
)

var Guesses []string
var Results [][]int
var UsedLetters map[rune]int
var Answer string

func init() {

	rand.Seed(int64(time.Now().Nanosecond()))
}

// Starts a new game.
func NewGame() {

	UsedLetters = make(map[rune]int)
	Answer = selectWord()
}

// Attempt to make a guess with the given string.
// If the guess string is invalid, an error is returned.
func MakeGuess(guess string) error {

	result, err := compareRunes(convertToRunes(guess), convertToRunes(Answer), UsedLetters)

	if err == nil {

		Guesses = append(Guesses, guess)
		Results = append(Results, result)
	}

	return err
}

// Returns true if the player has guessed the correct word.
func HasWon(guess string) bool {

	return guess == Answer
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
			result[i] = CorrectPosition
			usedLetters[r] = CorrectPosition
		}
	}

	// Check for characters in the wrong position
	for i, r := range guess {

		// Skip runes already checked
		if result[i] == CorrectPosition {

			continue
		}

		occurences[r] = occurences[r] + 1
		answerIndices := getRuneIndices(answer, r)

		// Ignore extra occurences of a rune
		if occurences[r] <= len(answerIndices) {

			result[i] = WrongPosition

			// Don't overwrite correct position status on letters
			if usedLetters[r] != CorrectPosition {

				usedLetters[r] = WrongPosition
			}
		}

		// Mark any letters that aren't in the word
		if usedLetters[r] != WrongPosition && usedLetters[r] != CorrectPosition {

			usedLetters[r] = NotInWord
		}
	}

	return result, nil
}

// Selects a random word from the list of answer words.
func selectWord() string {

	return Words[rand.Intn(len(Words))]
}
