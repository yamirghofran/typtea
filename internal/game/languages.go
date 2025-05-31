package game

import (
	"embed"
	"encoding/json"
	"fmt"
	"io/fs"
	"path/filepath"
	"strings"
)

//go:embed data/*.json
var embeddedLanguages embed.FS

// LanguageData represents the structure of the language JSON files
type LanguageData struct {
	Name  string   `json:"name"`
	Words []string `json:"words"`
}

// LanguageManager manages loading and caching of language data
type LanguageManager struct {
	loadedLanguages    map[string][]string
	availableLanguages []string
}

// NewLanguageManager initializes a new LanguageManager and scans for available languages
func NewLanguageManager() *LanguageManager {
	lm := &LanguageManager{
		loadedLanguages: make(map[string][]string),
	}
	lm.scanAvailableLanguages()
	return lm
}

// scanAvailableLanguages scans the embedded filesystem for available language files
func (lm *LanguageManager) scanAvailableLanguages() {
	entries, err := fs.ReadDir(embeddedLanguages, "data")
	if err != nil {
		return
	}

	for _, entry := range entries {
		if filepath.Ext(entry.Name()) == ".json" {
			lang := strings.TrimSuffix(entry.Name(), ".json")
			lm.availableLanguages = append(lm.availableLanguages, lang)
		}
	}
}

// LoadLanguage loads the specified language from embedded files and caches it
func (lm *LanguageManager) LoadLanguage(langCode string) ([]string, error) {
	// Check if already loaded
	if words, exists := lm.loadedLanguages[langCode]; exists {
		return words, nil
	}

	// Load from embedded files
	filename := fmt.Sprintf("data/%s.json", langCode)

	data, err := embeddedLanguages.ReadFile(filename)
	if err != nil {
		// Fallback to English if language not found
		if langCode != "en" {
			fmt.Printf("Language '%s' not found, falling back to English\n", langCode)
			return lm.LoadLanguage("en")
		}
		return nil, fmt.Errorf("could not load language data for '%s': %v", langCode, err)
	}

	var langData LanguageData
	if err := json.Unmarshal(data, &langData); err != nil {
		return nil, fmt.Errorf("could not parse language data for '%s': %v", langCode, err)
	}

	// Cache the loaded language
	lm.loadedLanguages[langCode] = langData.Words
	return langData.Words, nil
}

// GetAvailableLanguages returns a list of all available language codes
func (lm *LanguageManager) GetAvailableLanguages() []string {
	return lm.availableLanguages
}

// IsLanguageAvailable checks if a language is available in the manager
func (lm *LanguageManager) IsLanguageAvailable(langCode string) bool {
	for _, lang := range lm.availableLanguages {
		if lang == langCode {
			return true
		}
	}
	return false
}
