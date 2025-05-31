package tui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

// View renders the current state of the Model as a string for display
func (m Model) View() string {
	if m.showResults {
		return m.renderResults()
	}

	var sections []string

	timer := m.renderTimer()
	sections = append(sections, timer)

	textDisplay := m.renderText()
	sections = append(sections, textDisplay)

	content := lipgloss.JoinVertical(lipgloss.Left, sections...)

	return lipgloss.Place(
		m.width, m.height,
		lipgloss.Center, lipgloss.Center,
		content,
	)
}

// renderTimer formats the remaining time for display
func (m Model) renderTimer() string {
	remaining := m.game.GetRemainingTime()
	return timeStyle.Render(fmt.Sprintf("%d", remaining))
}

// renderText formats the text display with appropriate styles for typed, current, untyped characters
func (m Model) renderText() string {
	displayText := m.game.GetDisplayText()
	userPos := m.game.CurrentPos

	var rendered strings.Builder

	for i, char := range displayText {
		var styledChar string

		switch {
		case i < userPos:
			// Already typed
			if m.game.Errors[m.game.GlobalPos-(userPos-i)] {
				styledChar = incorrectCharStyle.Render(string(char))
			} else {
				styledChar = typedCharStyle.Render(string(char))
			}
		case i == userPos:
			// Current character
			styledChar = currentCharStyle.Render(string(char))
		default:
			// Not yet typed
			styledChar = untypedCharStyle.Render(string(char))
		}

		rendered.WriteString(styledChar)
	}

	// Format into lines
	content := rendered.String()
	lines := m.formatIntoLines(content)

	return textBoxStyle.Render(strings.Join(lines, "\n"))
}

// formatIntoLines formats the content into lines based on the game's display settings
func (m Model) formatIntoLines(plainContent string) []string {
	lines := m.game.DisplayLines

	maxLines := m.game.LinesPerView
	if len(lines) > maxLines {
		lines = lines[:maxLines]
	}

	var styledLines []string
	charIndex := 0

	for i, line := range lines {
		if i >= maxLines {
			break
		}

		if charIndex >= len(plainContent) {
			styledLines = append(styledLines, untypedCharStyle.Render(line))
			continue
		}

		var styledLine strings.Builder

		for _, char := range line {
			if charIndex < len(plainContent) {
				var styledChar string
				userPos := m.game.CurrentPos

				switch {
				case charIndex < userPos:
					if m.game.Errors[m.game.GlobalPos-(userPos-charIndex)] {
						styledChar = incorrectCharStyle.Render(string(char))
					} else {
						styledChar = typedCharStyle.Render(string(char))
					}
				case charIndex == userPos:
					styledChar = currentCharStyle.Render(string(char))
				default:
					styledChar = untypedCharStyle.Render(string(char))
				}

				styledLine.WriteString(styledChar)
				charIndex++
			} else {
				styledLine.WriteString(untypedCharStyle.Render(string(char)))
			}
		}

		styledLines = append(styledLines, styledLine.String())

		if charIndex < len(plainContent) && i < len(lines)-1 {
			charIndex++
		}
	}

	return styledLines
}

// renderResults formats the final results of the typing test for display
func (m Model) renderResults() string {
	stats := m.finalStats

	accSection := lipgloss.JoinVertical(
		lipgloss.Center,
		statLabelStyle.Render("acc"),
		statValueStyle.Render(fmt.Sprintf("%.0f%%", stats.Accuracy)),
	)

	wpmSection := lipgloss.JoinVertical(
		lipgloss.Center,
		statLabelStyle.Render("wpm"),
		statValueStyle.Render(fmt.Sprintf("%.0f", stats.WPM)),
	)

	timeSection := lipgloss.JoinVertical(
		lipgloss.Center,
		statLabelStyle.Render("time"),
		statValueStyle.Render(fmt.Sprintf("%.0fs", stats.TimeElapsed.Seconds())),
	)

	languageSection := lipgloss.JoinVertical(
		lipgloss.Center,
		statLabelStyle.Render("lang"),
		statValueStyle.Render(m.language),
	)

	// Arrange stats horizontally
	statsRow := lipgloss.JoinHorizontal(
		lipgloss.Top,
		accSection,
		strings.Repeat(" ", 5),
		wpmSection,
		strings.Repeat(" ", 5),
		timeSection,
		strings.Repeat(" ", 5),
		languageSection,
	)

	instructions := restartInstructionStyle.Render("Press Enter to restart • Esc to quit")

	// Results layout
	resultsContent := lipgloss.JoinVertical(
		lipgloss.Center,
		"",
		statsRow,
		"",
		instructions,
	)

	return lipgloss.Place(
		m.width, m.height,
		lipgloss.Center, lipgloss.Center,
		resultsContainerStyle.Render(resultsContent),
	)
}
