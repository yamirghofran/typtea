package tui

import (
	"github.com/charmbracelet/lipgloss"
)

// Styles for the TUI
var (
	timeStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("12")).
			Bold(true).
			MarginLeft(6)

	textBoxStyle = lipgloss.NewStyle().
			Padding(1, 3).
			Width(60).
			Height(6).
			Align(lipgloss.Left).
			MarginLeft(3)

	typedCharStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("15")).
			Bold(true)

	incorrectCharStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("9")).
				Bold(true).
				Underline(true)

	currentCharStyle = lipgloss.NewStyle().
				Background(lipgloss.Color("15")).
				Foreground(lipgloss.Color("0")).
				Bold(true)

	untypedCharStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("8"))

	statLabelStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("8")).
			Align(lipgloss.Center)

	statValueStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("15")).
			Bold(true).
			Align(lipgloss.Center)

	resultsContainerStyle = lipgloss.NewStyle().
				Padding(3, 5).
				Align(lipgloss.Center)

	restartInstructionStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("8")).
				Align(lipgloss.Center)
)
