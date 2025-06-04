package game

import (
	"math/rand"
	"sort"
	"strings"
	"time"
)

var languageManager *LanguageManager
var currentLanguageWords []string
var weights []int
var cumulativeWeights []int
var currentLanguageCode string

// init initializes the language manager and sets the default language to "en"
func init() {
	languageManager = NewLanguageManager()
	err := SetLanguage("en")
	if err != nil {
		panic("failed to initialize default language: " + err.Error())
	}
}

// SetLanguage sets the current language for the game and loads the corresponding words
func SetLanguage(langCode string) error {
	words, err := languageManager.LoadLanguage(langCode)
	if err != nil {
		return err
	}

	currentLanguageWords = words
	currentLanguageCode = langCode

	// Only calculate weights for English
	if langCode == "en" {
		calculateWeights()
	} else {
		weights = nil
		cumulativeWeights = nil
	}

	return nil
}

// calculateWeights calculates the weights for the words based on their rank
func calculateWeights() {
	if len(currentLanguageWords) == 0 {
		return
	}

	// Initialize weights inversely proportional to rank
	weights = make([]int, len(currentLanguageWords))
	for i := range currentLanguageWords {
		weights[i] = len(currentLanguageWords) - i
	}

	// Calculate cumulative weights for binary search
	cumulativeWeights = make([]int, len(weights))
	cumSum := 0
	for i, w := range weights {
		cumSum += w
		cumulativeWeights[i] = cumSum
	}
}

// findWordIndex uses binary search to find the index of the word based on the random number r
func findWordIndex(r int) int {
	if len(cumulativeWeights) == 0 {
		return 0
	}

	return sort.Search(len(cumulativeWeights), func(i int) bool {
		return cumulativeWeights[i] >= r
	})
}

// GenerateWords generates a slice of words based on the current language and the specified count
func GenerateWords(count int) []string {
	if len(currentLanguageWords) == 0 {
		// Fallback to English
		if err := SetLanguage("en"); err != nil {
			panic("failed to load fallback language: " + err.Error())
		}
	}

	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	words := make([]string, count)

	// Use weighted selection only for English
	if currentLanguageCode == "en" && len(cumulativeWeights) > 0 {
		maxWeight := cumulativeWeights[len(cumulativeWeights)-1]

		for i := range words {
			r := rng.Intn(maxWeight) + 1 // random in range [1, maxWeight]
			idx := findWordIndex(r)
			words[i] = currentLanguageWords[idx]
		}
		return words
	}

	// For all other languages, use simple random selection
	for i := range words {
		words[i] = currentLanguageWords[rng.Intn(len(currentLanguageWords))]
	}

	return words
}

// GenerateText generates a string of words joined by spaces
func GenerateText(words []string) string {
	return strings.Join(words, " ")
}
