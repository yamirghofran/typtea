package cmd

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/ashish0kumar/typtea/internal/game"
	"github.com/ashish0kumar/typtea/internal/tui"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

var (
	duration int    // Duration of the typing test in seconds
	language string // Language for the typing test, default is "en"
	listLangs bool  // Flag to list all available languages
)

// startCmd represents the start command for the typing test
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start a typing test",
	Long:  "Start a new typing test session with customizable duration and language",
	Example: `  typtea start --duration 60 --lang python
  typtea start -d 30 -l javascript
  typtea start --lang go
  typtea start --list-langs`,
	RunE: runTypingTest,
}

// init function initializes the start command and its flags
func init() {
	startCmd.Flags().IntVarP(&duration, "duration", "d", 30, "Test duration in seconds")
	startCmd.Flags().StringVarP(&language, "lang", "l", "en", "Language for typing test")
	startCmd.Flags().BoolVar(&listLangs, "list-langs", false, "List all available languages")
}

// runTypingTest is the main function that runs the typing test
func runTypingTest(cmd *cobra.Command, args []string) error {

	// Initialize the language manager
	langManager := game.NewLanguageManager()

	// If --list-langs is set, print available languages and exit
	if listLangs {
		fmt.Println("Available languages:")
		for _, lang := range langManager.GetAvailableLanguages() {
			fmt.Printf("  %s\n", lang)
		}
		return nil
	}

	// Validate duration
	if duration < 10 || duration > 300 {
		return fmt.Errorf("duration must be between 10 and 300 seconds (e.g., --duration 60)")
	}

	// Validate language
	if !langManager.IsLanguageAvailable(language) {
		available := langManager.GetAvailableLanguages()
		fmt.Fprintf(os.Stderr, "Error: Language '%s' not available.\n", language)
		fmt.Fprintf(os.Stderr, "Available languages: %s\n", strings.Join(available, ", "))
		return fmt.Errorf("invalid language: %s", language)
	}

	// Create a new typing test model
	model, err := tui.NewModel(duration, language)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating typing test: %v\n", err)
		os.Exit(1)
	}

	// Start the TUI program
	p := tea.NewProgram(model, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}

	return nil
}
