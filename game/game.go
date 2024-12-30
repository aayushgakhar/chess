package game

import (
	"chess/myGame"
	"strconv"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	dt "github.com/dylhunn/dragontoothmg"
)

type Game struct {
	board      *dt.Board
	moves      []dt.Move
	pieceMoves []dt.Move
	selected   string
	flipped    bool
	buffer     string
	message    string
}

func NewGame() *Game {
	return NewGameWithPosition(dt.Startpos)
}

func NewGameWithPosition(fen string) *Game {
	g := Game{flipped: false, message: "White to move"}
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
	row = max(0, min(7, row))
	col = max(0, min(7, col))
	if flipped {
		col = 7 - col
	} else {
		row = 7 - row
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

func (g *Game) setLegalSelected() {
	LegalMoves := []dt.Move{}
	if g.selected == "" {
		g.pieceMoves = LegalMoves
		return
	}
	for _, move := range g.moves {
		if strings.HasPrefix(move.String(), g.selected) {
			LegalMoves = append(LegalMoves, move)
		}
	}
	g.pieceMoves = LegalMoves
}

func (g *Game) isLegal(moves []dt.Move,destination string)bool {
	for _,move := range moves {
		if strings.HasSuffix(move.String(), destination){
			return true
		}
		if move.Promote() > 1 && strings.HasSuffix(move.String(), destination+"q"){
			return true
		}
	}
	return false
}

func (g *Game) Select(square string) (tea.Model, tea.Cmd) {
	if g.selected != "" {
		from := g.selected
		to := square
		for _, move := range g.pieceMoves {
			if move.String() == from+to || move.Promote() > 1 && move.String() == from+to+"q" {
				g.board.Apply(move)
				g.moves = g.board.GenerateLegalMoves()
				check := g.board.OurKingInCheck()
				checkmate := check && len(g.moves) == 0
				if checkmate {
					g.message = "Checkmate!"
				} else if check {
					g.message = "Check!"
				} else if g.board.Wtomove {
					g.message = "White to move"
				} else {
					g.message = "Black to move"
				}
				g.Deselect()
				return g, nil
			}
		}
		g.selected = square
	} else {
		g.selected = square
	}
	g.setLegalSelected()
	return g, nil
}

func (g *Game) Deselect() (tea.Model, tea.Cmd) {
	g.selected = ""
	g.pieceMoves = []dt.Move{}
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
	grid := Grid(g.board.ToFen())
	var s strings.Builder
	s.WriteString(Top)
	s.WriteString("\n")
	whiteTurn := g.board.Wtomove
	for ind, i := range rows {
		s.WriteString(strconv.Itoa(i + 1))
		s.WriteString(Sep)
		for _, j := range cols {
			current := myGame.CoordsToNotation(i, j)
			p := Piece(grid[i][j])
			display := p.Display()
			correct := p.IsWhite() == whiteTurn

			if current == g.selected {
				if correct {
					display = Cyan(display)
				}else{
					display = Red(display)
				}
			}
			if g.isLegal(g.pieceMoves,current){
				if p.IsEmpty(){
					display = "."
				}
				display = Magenta(display)
			}
			s.WriteString(" ")
			s.WriteString(display)
			s.WriteString(Sep)
		}
		s.WriteString("\n")
		if ind < 7 {
			s.WriteString(RowSep)
			s.WriteString("\n")
		}
	}
	s.WriteString(Bottom)
	s.WriteString("\n")
	s.WriteString("   ")
	for _, i := range cols {
		s.WriteString(" ")
		s.WriteString(string(rune('a' + i)))
		s.WriteString("  ")
	}
	s.WriteString("\n")
	s.WriteString(g.message)
	return s.String()
}
