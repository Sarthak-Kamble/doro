package ui

import "github.com/charmbracelet/lipgloss"

var (
	titleStyle = lipgloss.NewStyle().
			Bold(true).
			Align(lipgloss.Center).PaddingTop(2)

	timerStyle = lipgloss.NewStyle().
			Bold(true).
			Align(lipgloss.Center)

	statusStyle = lipgloss.NewStyle().
			Italic(true).
			Align(lipgloss.Center)

	helpStyle = lipgloss.NewStyle().
			Align(lipgloss.Center)
)
