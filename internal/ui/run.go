package ui

import (
	"doro/internal/config"

	tea "github.com/charmbracelet/bubbletea"
)

func Run(cfg config.Config) error {

	p := tea.NewProgram(
		NewModel(cfg),
	)

	_, err := p.Run()

	return err
}
