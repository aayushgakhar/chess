package game

import (
	"chess/myGame"
	"fmt"
	"strconv"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	dt "github.com/dylhunn/dragontoothmg"
)



type Game struct {
	board *dt.Board
	moves []dt.Move
	pieceMoves []dt.Move
	selected string
	flipped bool
	buffer string
}

func NewGame() *Game {
	return NewGameWithPosition(dt.Startpos)
}

func NewGameWithPosition(fen string) *Game {
	g := Game{}
	if !IsValidFen(fen) {
		fen = dt.Startpos
	}
	board := dt.ParseFen(fen)
	g.board = &board
	g.moves = g.board.GenerateLegalMoves()
	return &g
}

func (g *Game) Init() tea.Cmd {
	return nil
}

func Cell(x, y int, flipped bool) string {
	row := (y - 1) / 2
	col := (x - 2) / 4
	if flipped {
		row = 7 - row
		col = 7 - col
	}
	return myGame.CoordsToNotation(row, col)
}

func (g *Game) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.MouseMsg:
		if msg.Action != tea.MouseActionRelease || msg.Button != tea.MouseButtonLeft {
			return g, nil
		}

		// Find the square the user clicked on, this will either be our square
		// square for our piece or the destination square for a move if a piece is
		// already square and that destination square completes a legal move
		square := Cell(msg.X, msg.Y, g.flipped)
		return g.Select(square)
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return g, tea.Quit
		case "ctrl+f":
			g.flipped = !g.flipped
		case "a", "b", "c", "d", "e", "f", "g", "h":
			g.buffer = msg.String()
		case "1", "2", "3", "4", "5", "6", "7", "8":
			var move string
			if g.buffer != "" {
				move = g.buffer + msg.String()
				g.buffer = ""
			}
			return g.Select(move)
		case "esc":
			return g.Deselect()
		}
	}
	// case MoveMsg:
	// 	g.selected = msg.From
	// 	g.pieceMoves = moves.LegalSelected(g.moves, g.selected)
	// 	return g.Select(msg.To)
	// } 

	return g, nil
}

func (g *Game) Select(square string) (tea.Model, tea.Cmd) {
	fmt.Println("Selecting", square)
	return g, nil
}

func (g *Game) Deselect() (tea.Model, tea.Cmd) {
	fmt.Println("Deselecting")
	return g, nil
}

const (
	Top    = "  ┌───┬───┬───┬───┬───┬───┬───┬───┐"
	Bottom = "  └───┴───┴───┴───┴───┴───┴───┴───┘"
	RowSep = "  ├───┼───┼───┼───┼───┼───┼───┼───┤"
	Sep    = " │"
)


func (g *Game) View() string {
	rows := []int{7, 6, 5, 4, 3, 2, 1, 0}
	cols := []int{0, 1, 2, 3, 4, 5, 6, 7}
	if g.flipped {
		rows = []int{0, 1, 2, 3, 4, 5, 6, 7}
		cols = []int{7, 6, 5, 4, 3, 2, 1, 0}
	}
	var s strings.Builder
	// fmt.Println(Top)
	s.WriteString(Top)
	s.WriteString("\n")
	for ind, i := range rows {
		// fmt.Print(i + 1)
		s.WriteString(strconv.Itoa(i + 1))
		// fmt.Print(Sep)
		s.WriteString(Sep)
		for _, j := range cols {
			// if g.Cells[i][j].Piece != nil {
			// 	// fmt.Printf(" %s%s", b.Cells[i][j].Piece.Icon, Sep)
			// 	s.WriteString(" ")
			// 	s.WriteString(b.Cells[i][j].Piece.Icon)
			// 	s.WriteString(Sep)
			// } else {
				// fmt.Printf(" .%s", Sep)
				if j>-1{}
				s.WriteString(" .")
				s.WriteString(Sep)
			// }
		}
		// fmt.Println()
		s.WriteString("\n")
		if ind < 7 {
			// fmt.Println(RowSep)
			s.WriteString(RowSep)
			s.WriteString("\n")
		}
	}
	// fmt.Println(Bottom)
	s.WriteString(Bottom)
	s.WriteString("\n")
	// fmt.Print("   ")
	s.WriteString("   ")
	for _, i := range cols {
		// fmt.Printf(" %c  ", 'a'+i)
		s.WriteString(" ")
		s.WriteString(string('a'+i))
		s.WriteString("  ")
	}
	// fmt.Println()
	s.WriteString("\n")
	return s.String()
}



