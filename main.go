package main

import (
	"chess/game"
	"chess/myGame"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	dt "github.com/dylhunn/dragontoothmg"
)

func main() {
	args := os.Args

	my := false
	if len(args) > 1 {
		my = args[1] == "1"
	}
	if my {
		myGame.StartGame()
	} else {
		p := tea.NewProgram(
			game.NewGameWithPosition(dt.Startpos),
			tea.WithAltScreen(),
			tea.WithMouseCellMotion(),
		)
		_, err := p.Run()
		if err != nil {
			panic(err)
		}
	}
}
