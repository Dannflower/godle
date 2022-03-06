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

// Returns true if both maps contain identical key/value pairs.
func usedLettersEqual(actual map[rune]int, expected map[rune]int) bool {

	if len(actual) != len(expected) {
		return false
	}

	for key, value := range actual {
		if value != expected[key] {
			return false
		}
	}

	return true
}

// Returns true if both slices contain the same runes
func runesEqual(actual []rune, expected []rune) bool {

	if len(actual) != len(expected) {
		return false
	}

	for i, r := range actual {
		if r != expected[i] {
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
	expectedUsedLetters := map[rune]int{
		'V': 1,
		'A': 2,
		'L': 1,
		'I': 2,
		'D': 0,
	}
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

	if !usedLettersEqual(UsedLetters, expectedUsedLetters) {
		t.Fatalf("MakeGuess(%s) did not add correct used letters. Expected: %v, Actual: %v.", guess, expectedUsedLetters, UsedLetters)
	}

	// Second valid guess
	guess = "folds"
	expectedResult = []int{0, 0, 1, 0, 0}
	expectedUsedLetters = map[rune]int{
		'V': 1,
		'A': 2,
		'L': 1,
		'I': 2,
		'D': 0,
		'F': 0,
		'O': 0,
		'S': 0,
	}
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

	if !usedLettersEqual(UsedLetters, expectedUsedLetters) {
		t.Fatalf("MakeGuess(%s) did not add correct used letters. Expected: %v, Actual: %v.", guess, expectedUsedLetters, UsedLetters)
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

	if !usedLettersEqual(UsedLetters, expectedUsedLetters) {
		t.Fatalf("MakeGuess(%s) did not add correct used letters. Expected: %v, Actual: %v.", guess, expectedUsedLetters, UsedLetters)
	}
}

func TestCompareRunes(t *testing.T) {

	// Test too many of the same letter
	// Should mark first 'a' as in word, second as not in word, and third in correct place.
	NewGame()
	answer := []rune{'b', 'b', 'a', 'a'}
	guess := []rune{'a', 'a', 'a', 'c'}
	expectedResult := []int{2, 0, 1, 0}
	expectedUsedLetters := map[rune]int{
		'a': 1,
		'c': 0,
	}

	actualResult, err := compareRunes(guess, answer)

	if err != nil {
		t.Fatalf("compareRunes(%v, %v) returned an error when it shouldn't. Error: %v", guess, answer, err)
	}

	if !resultsEqual(actualResult, expectedResult) {
		t.Fatalf("compareRunes(%v, %v) return incorrect results. Expected %v, Actual: %v", guess, answer, expectedResult, actualResult)
	}

	if !usedLettersEqual(UsedLetters, expectedUsedLetters) {
		t.Fatalf("compareRunes(%v, %v) did not correctly mark used letters. Expected: %v, Actual: %v", guess, answer, expectedUsedLetters, UsedLetters)
	}

	// Test one each: in word, wrong spot, not in word
	NewGame()
	answer = []rune{'a', 'b', 'c'}
	guess = []rune{'a', 'c', 'z'}
	expectedResult = []int{1, 2, 0}
	expectedUsedLetters = map[rune]int{
		'a': 1,
		'c': 2,
		'z': 0,
	}

	actualResult, err = compareRunes(guess, answer)

	if err != nil {
		t.Fatalf("compareRunes(%v, %v) returned an error when it shouldn't. Error: %v", guess, answer, err)
	}

	if !resultsEqual(actualResult, expectedResult) {
		t.Fatalf("compareRunes(%v, %v) return incorrect results. Expected %v, Actual: %v", guess, answer, expectedResult, actualResult)
	}

	if !usedLettersEqual(UsedLetters, expectedUsedLetters) {
		t.Fatalf("compareRunes(%v, %v) did not correctly mark used letters. Expected: %v, Actual: %v", guess, answer, expectedUsedLetters, UsedLetters)
	}

	// Test incompatible lengths
	NewGame()
	answer = []rune{'a', 'b', 'c'}
	guess = []rune{'a', 'b', 'c', 'd'}

	actualResult, err = compareRunes(guess, answer)

	if err == nil {
		t.Fatalf("compareRunes(%v, %v) did not return an error when comparing slices of different lengths.", guess, answer)
	}

	if actualResult != nil {
		t.Fatalf("compareRunes(%v, %v) returned a non-nil result when comparing slices of different lengths.", guess, answer)
	}
}

func TestSelectWord(t *testing.T) {

	NewGame()
	word := selectWord()

	if !isValidAnswerWord(word) {
		t.Fatalf("selectWord() returned %s, which is not a valid answer word.", word)
	}
}

func TestHasWon(t *testing.T) {

	// Correct guess
	NewGame()
	guess := Answer

	if !HasWon(guess) {
		t.Fatalf("HasWon(%s) returned false when answer was '%s'.", guess, Answer)
	}

	// Incorrect guess
	NewGame()
	Answer = "words"
	guess = "fjord"

	if HasWon(guess) {
		t.Fatalf("HasWon(%s) returned true when answer was '%s'.", guess, Answer)
	}
}

func TestGetRuneIndices(t *testing.T) {

	// Rune present
	NewGame()
	runes := []rune{'a', 'b', 'c', 'a'}
	r := 'a'
	expectedResult := []int{0, 3}
	actualResult := getRuneIndices(runes, r)

	if !resultsEqual(actualResult, expectedResult) {
		t.Fatalf("getRuneIndices(%v, %v) returned incorrect results. Expected: %v, Actual: %v", runes, r, expectedResult, actualResult)
	}

	// Rune not present
	NewGame()
	runes = []rune{'z', 'b', 'c', 'z'}
	r = 'a'
	expectedResult = []int{}
	actualResult = getRuneIndices(runes, r)

	if !resultsEqual(actualResult, expectedResult) {
		t.Fatalf("getRuneIndices(%v, %v) returned incorrect results. Expected: %v, Actual: %v", runes, r, expectedResult, actualResult)
	}
}

func TestConvertToRunes(t *testing.T) {

	// Non-empty string
	input := "string"
	expected := []rune{'S', 'T', 'R', 'I', 'N', 'G'}
	actual := convertToRunes(input)

	if !runesEqual(actual, expected) {
		t.Fatalf("convertToRunes(%s) returned %v, when %v was expected.", input, actual, expected)
	}

	// Empty string
	input = ""
	expected = []rune{}
	actual = convertToRunes(input)

	if !runesEqual(actual, expected) {
		t.Fatalf("convertToRunes(%s) returned %v, when %v was expected.", input, actual, expected)
	}
}

func TestIsDuplicateGuess(t *testing.T) {

	// New guess
	NewGame()
	guess := "piety"
	if isDuplicateGuess(guess) {
		t.Fatalf("isDuplicateGuess(%s) returned true on a non-duplicate guess.", guess)
	}

	// Duplicate guess
	MakeGuess(guess)
	if !isDuplicateGuess(guess) {
		t.Fatalf("isDuplicateGuess(%s) returned false on a duplicate guess.", guess)
	}
}

func TestIsValidWord(t *testing.T) {

	// Valid word
	word := "piety"
	if !isValidWord(word) {
		t.Fatalf("isValidWord(%s) returned false for a valid word.", word)
	}

	// Invalid word
	word = "bbbbb"
	if isValidWord(word) {
		t.Fatalf("isValidWord(%s) returned true for an invalid word.", word)
	}
}
