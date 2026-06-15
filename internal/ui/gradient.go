package ui

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

var gradientStops = [][3]int{
	{0, 121, 145},     // ##007991
	{120, 255, 214}, // #1CB5E0
}


func interpolateColor(position float64) string {

	start := gradientStops[0]
	end := gradientStops[1]


	r := int(
		float64(start[0])+
			(float64(end[0])-float64(start[0]))*position,
	)

	g := int(
		float64(start[1])+
			(float64(end[1])-float64(start[1]))*position,
	)

	b := int(
		float64(start[2])+
			(float64(end[2])-float64(start[2]))*position,
	)


	return fmt.Sprintf(
		"#%02x%02x%02x",
		r,
		g,
		b,
	)
}


func gradientBlock(position float64) string {

	return lipgloss.NewStyle().
		Foreground(
			lipgloss.Color(
				interpolateColor(position),
			),
		).
		Render("█")
}