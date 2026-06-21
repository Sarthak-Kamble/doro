package ui

import (
	"doro/internal/app"
	"doro/internal/config"
	"fmt"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type tickMsg struct{}

type Model struct {
	doro *app.Doro
}


func NewModel(cfg config.Config) Model {
    return Model{
        doro: app.NewDoro(cfg),
    }
}

func tickCmd() tea.Cmd {

	return tea.Tick(
		time.Second,
		func(time.Time) tea.Msg {
			return tickMsg{}
		},
	)
}

func (m Model) Init() tea.Cmd {
	return tickCmd()
}

func (m Model) Update(
	msg tea.Msg,
) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {

	case tickMsg:

		m.doro.Tick()
		return m, tickCmd()

	case tea.KeyMsg:

		switch msg.String() {

		case " ":
			m.doro.IsPaused()

		case "r":
			m.doro.Reset()

		case "q", "ctrl+c":
			// return m, tea.Quit
			return m, func() tea.Msg {
				return tea.Quit()
			}
		}

	}

	return m, nil
}

func (m Model) View() string {

	phase := m.doro.CurrentPhase()

	session := m.doro.Session()

	status := "Running"

	if m.doro.IsPaused() {
		status = "Paused"
	}

	if m.doro.Timer.Remaining <= 0 {
		status = "Completed"
	}

	phaseText := fmt.Sprintf(
		"%s | Session %d/4",
		phase.String(),
		session,
	)

	title := titleStyle.Render(
		"🍅 Doro",
	)

	timer := timerStyle.Render(
		m.doro.Timer.FormatRemaining(),
	)

	bar := timerStyle.Render(
		m.progressBar(),
	)

	statusText := statusStyle.Render(
		"Status: " + status,
	)

	help := helpStyle.Render(
		"SPACE: Pause | R: Reset | Q: Quit",
	)

	return lipgloss.JoinVertical(
		lipgloss.Center,
		title,
		"",
		timer,
		"",
		phaseText,
		"",
		bar,
		"",
		statusText,
		"",
		help,
	)
}

func (m Model) progress() float64 {

	total := m.doro.Timer.Duration.Seconds()

	remaining := m.doro.Timer.Remaining.Seconds()

	return remaining / total
}

func (m Model) progressBar() string {

	width := 35

	filled := int(
		m.progress() * float64(width),
	)

	var bar strings.Builder

	for i := 0; i < filled; i++ {

		position := float64(i) / float64(width)

		bar.WriteString(
			gradientBlock(position),
		)
	}

	empty := lipgloss.NewStyle().
		Foreground(
			lipgloss.Color("#444444"),
		)

	for i := filled; i < width; i++ {

		bar.WriteString(
			empty.Render("░"),
		)
	}

	return bar.String()
}
