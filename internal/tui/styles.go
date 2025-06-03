package tui

import (
	"github.com/charmbracelet/lipgloss"
)

// Styles for the TUI
var (
	timeStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("12")).
			Bold(true).
			MarginLeft(8)

	textBoxStyle = lipgloss.NewStyle().
			Padding(1, 3).
			Width(60).
			Height(6).
			Align(lipgloss.Left).
			MarginLeft(5)

	boldStyle = lipgloss.NewStyle().
			Bold(true)

	mutedStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("8"))

	errorStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("9")).
			Bold(true).
			Underline(true)

	cursorStyle = lipgloss.NewStyle().
			Background(lipgloss.Color("15")).
			Foreground(lipgloss.Color("#000")).
			Bold(true)

	resultsContainerStyle = lipgloss.NewStyle().
				Padding(3, 5).
				Align(lipgloss.Left)
)
