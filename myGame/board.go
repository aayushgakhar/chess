package myGame

import (
	"errors"
	"fmt"
)

type Board struct {
	Cells          [][]Cell
	Pieces         []Piece
	CapturedPieces []Piece
	CurrentPlayer  string
	Moves          []Move
}

func NewBoard() *Board {
	b := Board{CurrentPlayer: White}
	b.Cells = make([][]Cell, 8)
	for i := range b.Cells {
		b.Cells[i] = make([]Cell, 8)
	}
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			b.Cells[i][j] = Cell{Row: i, Column: j}
		}
	}
	b.initialisePieces()
	return &b
}

func (b *Board) ChangeTurn() {
	if b.CurrentPlayer == White {
		b.CurrentPlayer = Black
	} else {
		b.CurrentPlayer = White
	}
}

func (b *Board) initialisePieces() {
	b.AddPiece(*NewPiece(White, ROOK), 0, 0)
	b.AddPiece(*NewPiece(White, KNIGHT), 0, 1)
	b.AddPiece(*NewPiece(White, BISHOP), 0, 2)
	b.AddPiece(*NewPiece(White, QUEEN), 0, 3)
	b.AddPiece(*NewPiece(White, KING), 0, 4)
	b.AddPiece(*NewPiece(White, BISHOP), 0, 5)
	b.AddPiece(*NewPiece(White, KNIGHT), 0, 6)
	b.AddPiece(*NewPiece(White, ROOK), 0, 7)
	for i := 0; i < 8; i++ {
		b.AddPiece(*NewPiece(White, PAWN), 1, i)
	}
	b.AddPiece(*NewPiece(Black, ROOK), 7, 0)
	b.AddPiece(*NewPiece(Black, KNIGHT), 7, 1)
	b.AddPiece(*NewPiece(Black, BISHOP), 7, 2)
	b.AddPiece(*NewPiece(Black, QUEEN), 7, 3)
	b.AddPiece(*NewPiece(Black, KING), 7, 4)
	b.AddPiece(*NewPiece(Black, BISHOP), 7, 5)
	b.AddPiece(*NewPiece(Black, KNIGHT), 7, 6)
	b.AddPiece(*NewPiece(Black, ROOK), 7, 7)
	for i := 0; i < 8; i++ {
		b.AddPiece(*NewPiece(Black, PAWN), 6, i)
	}
}

func (b *Board) GetCell(row, column int) *Cell {
	return &b.Cells[row][column]
}

func (b *Board) AddPiece(p Piece, row, column int) {
	p.Cell = &b.Cells[row][column]
	b.Cells[row][column].Piece = &p
	b.Pieces = append(b.Pieces, p)
}

// const (
// 	Top    = "  +------------------------+"
// 	Bottom = "  +------------------------+"
// 	Sep    = "|"
// )

const (
	Top    = "  ┌───┬───┬───┬───┬───┬───┬───┬───┐"
	Bottom = "  └───┴───┴───┴───┴───┴───┴───┴───┘"
	RowSep = "  ├───┼───┼───┼───┼───┼───┼───┼───┤"
	Sep    = " │"
)
//    ┌───┬───┬───┬───┬───┬───┬───┬───┐
//  8 │ ♖ │ ♘ │ ♗ │ ♕ │ ♔ │ ♗ │ ♘ │ ♖ │
//    ├───┼───┼───┼───┼───┼───┼───┼───┤
//  7 │ ♙ │ ♙ │ ♙ │ ♙ │ ♙ │ ♙ │ ♙ │ ♙ │
//    ├───┼───┼───┼───┼───┼───┼───┼───┤
//  6 │   │   │   │   │   │   │   │   │
//    ├───┼───┼───┼───┼───┼───┼───┼───┤
//  5 │   │   │   │   │   │   │   │   │
//    ├───┼───┼───┼───┼───┼───┼───┼───┤
//  4 │   │   │   │   │ . │   │   │   │
//    ├───┼───┼───┼───┼───┼───┼───┼───┤
//  3 │   │   │   │   │ . │   │   │   │
//    ├───┼───┼───┼───┼───┼───┼───┼───┤
//  2 │ ♟ │ ♟ │ ♟ │ ♟ │ ♟ │ ♟ │ ♟ │ ♟ │
//    ├───┼───┼───┼───┼───┼───┼───┼───┤
//  1 │ ♜ │ ♞ │ ♝ │ ♛ │ ♚ │ ♝ │ ♞ │ ♜ │
//    └───┴───┴───┴───┴───┴───┴───┴───┘
//      A   B   C   D   E   F   G   H

func (b *Board) PrintBoard(flipped bool) {
	rows := []int{7, 6, 5, 4, 3, 2, 1, 0}
	cols := []int{0, 1, 2, 3, 4, 5, 6, 7}
	if flipped {
		rows = []int{0, 1, 2, 3, 4, 5, 6, 7}
		cols = []int{7, 6, 5, 4, 3, 2, 1, 0}
	}
	fmt.Println(Top)
	for ind, i := range rows {
		fmt.Print(i + 1)
		fmt.Print( Sep)
		for _, j := range cols {
			if b.Cells[i][j].Piece != nil {
				fmt.Printf(" %s%s", b.Cells[i][j].Piece.Icon, Sep)
			} else {
				fmt.Printf(" .%s", Sep)
			}
		}
		fmt.Println()
		if ind < 7 {
			fmt.Println(RowSep)
		}
	}
	fmt.Println(Bottom)
	fmt.Print("   ")
	for _, i := range cols {
		fmt.Printf(" %c  ", 'a'+i)
	}
	fmt.Println()
	if len(b.CapturedPieces) > 0 {
		fmt.Println("Captured Pieces:")
		for _, p := range b.CapturedPieces {
			fmt.Printf("%s ", p.Icon)
		}
		fmt.Println()
	}
}

func (b *Board) MovePiece(p *Piece, row, column int) {
	p.Cell.Piece = nil
	p.Cell = &b.Cells[row][column]
	if b.Cells[row][column].Piece != nil {
		b.CapturedPieces = append(b.CapturedPieces, *b.Cells[row][column].Piece)
	}
	b.Cells[row][column].Piece = p
}

func (b *Board) IsValidMove(curr, new string) bool {
	r1, c1 := NotationToCoords(curr)
	if c1 < 0 || c1 > 7 || r1 < 0 || r1 > 7 {
		return false
	}
	p := b.GetCell(r1, c1).Piece
	if p == nil || p.Color != b.CurrentPlayer {
		fmt.Println("Invalid piece")
		return false
	}
	r2, c2 := NotationToCoords(new)
	if c2 < 0 || c2 > 7 || r2 < 0 || r2 > 7 {
		return false
	}
	return true
}

func (b *Board) MakeMove(curr, new string) error {
	if !b.IsValidMove(curr, new) {
		fmt.Println("Invalid move")
		return errors.New("Invalid move")
	}
	r1, c1 := NotationToCoords(curr)
	p := b.GetCell(r1, c1).Piece
	r2, c2 := NotationToCoords(new)
	b.MovePiece(p, r2, c2)
	b.Moves = append(b.Moves, Move{Piece: *p, From: b.GetCell(r1, c1), To: b.GetCell(r2, c2)})
	b.ChangeTurn()
	return nil
}

func (b *Board) GameOver() bool {
	return false
}

func NotationToCoords(notation string) (int, int) {
	return int(notation[1] - '1'), int(notation[0] - 'a')
}

func CoordsToNotation(row, col int) string {
	return string(rune('a'+col)) + string(rune('1'+row))
}
