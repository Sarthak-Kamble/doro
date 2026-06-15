package ui

import "github.com/charmbracelet/lipgloss"

var (
	titleStyle = lipgloss.NewStyle().
			Bold(true).
			Align(lipgloss.Center).PaddingTop(2).PaddingLeft(2).PaddingRight(2)

	timerStyle = lipgloss.NewStyle().
			Bold(true).
			Align(lipgloss.Center).PaddingLeft(2).PaddingRight(2)

	statusStyle = lipgloss.NewStyle().
			Italic(true).
			Align(lipgloss.Center).PaddingLeft(2).PaddingRight(2)

	helpStyle = lipgloss.NewStyle().
			Align(lipgloss.Center).PaddingLeft(2).PaddingRight(2)
)
