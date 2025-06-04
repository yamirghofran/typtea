package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	version     = "dev" // fallback to dev
	showVersion bool
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
	Run: func(cmd *cobra.Command, args []string) {
		// Show help if no subcommands or flags are provided
		cmd.Help()
	},
}

// versionCmd prints the current version of typtea
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show the version of typtea",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("typtea version", version)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// init function initializes the root command and adds subcommands and flags
func init() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true // Disable default completion command

	// Add --version flag with shorthand -v
	rootCmd.PersistentFlags().BoolVarP(&showVersion, "version", "v", false, "Show the version and exit")

	// Add your subcommands
	rootCmd.AddCommand(startCmd)
	rootCmd.AddCommand(versionCmd)

	// Check for version flag early and exit if set
	cobra.OnInitialize(func() {
		if showVersion {
			fmt.Println("typtea version", version)
			os.Exit(0)
		}
	})
}
