package logic

import (
	"strings"
	"testing"
)

// Returns true if the given answer word is in the
// list of valid answer words.
func isValidAnswerWord(answer string) bool {

	for _, word := range AnswerWords {

		if answer == word {
			return true
		}
	}

	return false
}

// Returns true if the two slices contain the same integers.
func resultsEqual(actual []int, expected []int) bool {

	if len(actual) != len(expected) {
		return false
	}

	for i, v := range actual {

		if v != expected[i] {
			return false
		}
	}

	return true
}

func TestHasWonNo(t *testing.T) {

	Answer = "aword"
	guess := "bword"
	expected := false
	result := HasWon(guess)

	if result {
		t.Fatalf("HasWon(%s) returned %v when Answer was %s, expected %v.", guess, result, Answer, expected)
	}
}

func TestHasWonYes(t *testing.T) {

	Answer = "aword"
	guess := "aword"
	expected := true
	result := HasWon(guess)

	if !result {
		t.Fatalf("HasWon(%s) returned %v when Answer was %s, expected %v.", guess, result, Answer, expected)
	}
}

func TestNewGame(t *testing.T) {

	// First new game
	NewGame()

	if !isValidAnswerWord(Answer) {
		t.Fatalf("NewGame() set Answer to %s, which is not in the list of answer words.", Answer)
	}

	if UsedLetters == nil {
		t.Fatal("NewGame() did not initialize UsedLetters.")
	}

	if Guesses != nil {
		t.Fatal("NewGame() did not clear Guesses.")
	}

	// New game after another has started
	MakeGuess("valid")
	NewGame()

	if !isValidAnswerWord(Answer) {
		t.Fatalf("NewGame() set Answer to %s, which is not in the list of answer words.", Answer)
	}

	if UsedLetters == nil {
		t.Fatal("NewGame() did not initialize UsedLetters.")
	}

	if Guesses != nil {
		t.Fatal("NewGame() did not clear Guesses.")
	}
}

func TestAnswersAreValidWords(t *testing.T) {

	for _, answer := range AnswerWords {

		isValid := false

		for _, validWord := range ValidWords {

			if strings.EqualFold(answer, validWord) {
				isValid = true
				break
			}
		}

		if !isValid {
			t.Fatalf("AnswersWords contains word '%s' which is not in ValidWords.", answer)
		}
	}
}

func TestMakeGuessInvalidWord(t *testing.T) {

	NewGame()

	// Invalid word, but right length
	guess := "aaaaa"
	result := MakeGuess(guess)

	if result == nil || result.Error() != "must be a valid word" {
		t.Fatalf("MakeGuess(%s) did not return an error for the invalid word.", guess)
	}

	if len(Guesses) > 0 {
		t.Fatalf("MakeGuess(%s) added invalid word to Guesses.", guess)
	}

	if len(Results) > 0 {
		t.Fatalf("MakeGuess(%s) added invalid word to Results.", guess)
	}

	// Too short
	guess = "aaaa"
	result = MakeGuess(guess)

	if result == nil || result.Error() != "must be a valid word" {
		t.Fatalf("MakeGuess(%s) did not return an error for the invalid word.", guess)
	}

	if len(Guesses) > 0 {
		t.Fatalf("MakeGuess(%s) added invalid word to Guesses.", guess)
	}

	if len(Results) > 0 {
		t.Fatalf("MakeGuess(%s) added invalid word to Results.", guess)
	}

	// Too long
	guess = "aaaaaa"
	result = MakeGuess(guess)

	if result == nil || result.Error() != "must be a valid word" {
		t.Fatalf("MakeGuess(%s) did not return an error for the invalid word.", guess)
	}

	if len(Guesses) > 0 {
		t.Fatalf("MakeGuess(%s) added invalid word to Guesses.", guess)
	}

	if len(Results) > 0 {
		t.Fatalf("MakeGuess(%s) added invalid word to Results.", guess)
	}

	// No word
	guess = ""
	result = MakeGuess(guess)

	if result == nil || result.Error() != "must be a valid word" {
		t.Fatalf("MakeGuess(%s) did not return an error for the invalid word.", guess)
	}

	if len(Guesses) > 0 {
		t.Fatalf("MakeGuess(%s) added invalid word to Guesses.", guess)
	}

	if len(Results) > 0 {
		t.Fatalf("MakeGuess(%s) added invalid word to Results.", guess)
	}
}

func TestMakeGuessValidWord(t *testing.T) {

	NewGame()

	// Choose a fixed answer to outputs are always the same
	Answer = "vilag"

	// First valid guess
	guess := "valid"
	expectedResult := []int{1, 2, 1, 2, 0}
	err := MakeGuess(guess)

	if err != nil {
		t.Fatalf("MakeGuess(%s) returned an error for valid word.", guess)
	}

	if len(Guesses) != 1 || !strings.EqualFold(Guesses[0], guess) {
		t.Fatalf("MakeGuess(%s) did not add valid word to Guesses.", guess)
	}

	if len(Results) != 1 || !resultsEqual(Results[0], expectedResult) {
		t.Fatalf("MakeGuess(%s) produced result %v, but expected %v.", guess, Results[0], expectedResult)
	}

	// Second valid guess
	guess = "folds"
	expectedResult = []int{0, 0, 1, 0, 0}
	err = MakeGuess(guess)

	if err != nil {
		t.Fatalf("MakeGuess(%s) returned an error for valid word.", guess)
	}

	if len(Guesses) != 2 || !strings.EqualFold(Guesses[1], guess) {
		t.Fatalf("MakeGuess(%s) did not add valid word to Guesses.", guess)
	}

	if len(Results) != 2 || !resultsEqual(Results[1], expectedResult) {
		t.Fatalf("MakeGuess(%s) produced result %v, but expected %v.", guess, Results[1], expectedResult)
	}

	// Duplicate guess
	err = MakeGuess(guess)

	if err == nil {
		t.Fatalf("MakeGuess(%s) allowed duplicate word.", guess)
	}

	if len(Guesses) > 2 {
		t.Fatalf("MakeGuess(%s) added duplicate word to Guesses.", guess)
	}

	if len(Results) > 2 {
		t.Fatalf("MakeGuess(%s) added duplicate word to Results.", guess)
	}
}
