package game

import lg "github.com/charmbracelet/lipgloss"

type colorFunc func(s ...string) string

func fg(color string) colorFunc {
	return lg.NewStyle().Foreground(lg.Color(color)).Render
}

var Cyan = fg("6")
var Faint = fg("8")
var Magenta = fg("5")
var Red = fg("1")

var Title = lg.NewStyle().Foreground(lg.Color("5")).Italic(true).Render