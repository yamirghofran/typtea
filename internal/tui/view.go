package tui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

const statGap = 5
const spacer = ""

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

	var rendered strings.Builder

	for i, char := range displayText {
		// Use helper to style character
		styledChar := m.styleChar(char, i)
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

		var styledLine strings.Builder

		lineRunes := []rune(line)

		for col := 0; col < len(lineRunes); col++ {
			if charIndex < len(plainContent) {
				styledChar := m.styleChar(lineRunes[col], charIndex)
				styledLine.WriteString(styledChar)
				charIndex++
			} else {
				styledLine.WriteString(mutedStyle.Render(string(lineRunes[col])))
			}
		}

		// Check if caret is on this line and positioned just beyond last char
		caretPos := m.game.CurrentPos
		if i == 0 && caretPos == len(lineRunes) {
			// Append caret style with a space or block to show cursor
			styledLine.WriteString(cursorStyle.Render(" "))
		}

		styledLines = append(styledLines, styledLine.String())

		// No extra space between lines
		// charIndex++ // This was adding an extra space
	}

	return styledLines
}

// styleChar determines the style of a character based on its position and error status
func (m Model) styleChar(char rune, index int) string {
	userPos := m.game.CurrentPos
	errorIndex := m.game.GlobalPos - (userPos - index)

	switch {
	case index < userPos:
		// Already typed
		if m.game.Errors != nil {
			if _, hasErr := m.game.Errors[errorIndex]; hasErr {
				return errorStyle.Render(string(char))
			}
		}
		return boldStyle.Render(string(char))
	case index == userPos:
		// Current character
		return cursorStyle.Render(string(char))
	default:
		// Not yet typed
		return mutedStyle.Render(string(char))
	}
}

// renderResults formats the final results of the typing test for display
func (m Model) renderResults() string {
	stats := m.finalStats

	accSection := lipgloss.JoinVertical(
		lipgloss.Right,
		mutedStyle.Render("acc"),
		boldStyle.Render(fmt.Sprintf("%.0f%%", stats.Accuracy)),
	)

	wpmSection := lipgloss.JoinVertical(
		lipgloss.Right,
		mutedStyle.Render("wpm"),
		boldStyle.Render(fmt.Sprintf("%.0f", stats.WPM)),
	)

	timeSection := lipgloss.JoinVertical(
		lipgloss.Right,
		mutedStyle.Render("time"),
		boldStyle.Render(fmt.Sprintf("%.0fs", stats.TimeElapsed.Seconds())),
	)

	languageSection := lipgloss.JoinVertical(
		lipgloss.Right,
		mutedStyle.Render("lang"),
		boldStyle.Render(m.language),
	)

	// Arrange stats horizontally
	statsRow := lipgloss.JoinHorizontal(
		lipgloss.Top,
		accSection,
		strings.Repeat(" ", statGap),
		wpmSection,
		strings.Repeat(" ", statGap),
		timeSection,
		strings.Repeat(" ", statGap),
		languageSection,
	)

	instructions := mutedStyle.Align(lipgloss.Center).Render("Press Enter to restart • Esc to quit")

	// Results layout
	resultsContent := lipgloss.JoinVertical(
		lipgloss.Center,
		spacer,
		statsRow,
		spacer,
		instructions,
	)

	return lipgloss.Place(
		m.width, m.height,
		lipgloss.Center, lipgloss.Center,
		resultsContainerStyle.Render(resultsContent),
	)
}
