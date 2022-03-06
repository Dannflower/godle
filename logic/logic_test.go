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

	NewGame()

	if !isValidAnswerWord(Answer) {
		t.Fatalf("NewGame() set Answer to %s, which is not in the list of answer words.", Answer)
	}

	if UsedLetters == nil {
		t.Fatal("NewGame() did not initialize UsedLetters.")
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
