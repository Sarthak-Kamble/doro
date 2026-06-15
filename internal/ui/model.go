package ui

import (
	"doro/internal/config"
	"doro/internal/timer"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type tickMsg struct{}

type Model struct {
	timer  *timer.Timer
	paused bool
}

func NewModel(cfg config.Config) Model {

	return Model{
		timer: timer.New(
			cfg.WorkDuration,
		),
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

		if !m.paused {
			m.timer.Tick()
		}

		return m, tickCmd()

	case tea.KeyMsg:

		switch msg.String() {

		case " ":
			m.paused = !m.paused

		case "r":
			m.timer.Reset()

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

	status := "Running"

	if m.paused {
		status = "Paused"
	}

	if m.timer.Remaining <= 0 {
		status = "Completed"
	}

	title := titleStyle.Render(
		"🍅 Doro",
	)

	timer := timerStyle.Render(
		m.timer.FormatRemaining(),
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
		bar,
		"",
		statusText,
		"",
		help,
	)
}

func (m Model) progress() float64 {

	total := m.timer.Duration.Seconds()

	remaining := m.timer.Remaining.Seconds()

	return remaining / total
}

// func (m Model) progressBar() string {

// 	width := 30

// 	filled := int(
// 		m.progress() * float64(width),
// 	)

// 	return strings.Repeat(
// 		"█",
// 		filled,
// 	) +
// 		strings.Repeat(
// 			"░",
// 			width-filled,
// 		)
// }

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
