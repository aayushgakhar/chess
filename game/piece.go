package game

import "strings"

type Piece string

var display = map[string]string{
	"":  " ",
	"B": "♝",
	"K": "♚",
	"N": "♞",
	"P": "♟",
	"Q": "♛",
	"R": "♜",
	"b": "♗",
	"k": "♔",
	"n": "♘",
	"p": "♙",
	"q": "♕",
	"r": "♖",
}

func (p *Piece) String() string {
	return string(*p);
}

func (p *Piece) Display() string {
	return display[p.String()];
}

func (p *Piece) IsWhite() bool {
	return strings.ToUpper(p.String()) == p.String()
}

func (p *Piece) IsEmpty() bool {
	return p.String() == ""
}

func (p *Piece) IsBlack() bool {
	return strings.ToLower(p.String()) == p.String()
}


