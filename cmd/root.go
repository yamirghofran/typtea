package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "typtea",
	Short: "A minimal typing speed test in your terminal",
	Long: `A terminal-based typing speed test application.
Supports multiple programming languages like Python, JavaScript, Go, and more.`,
	Example: `  typtea start --lang python
  typtea start --duration 30 --lang javascript
  typtea start --list-langs`,
}

// Execute adds all child commands to the root command and sets flags appropriately
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// init function initializes the root command and adds subcommands
func init() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true // Disable default completion command
	rootCmd.AddCommand(startCmd)
}
