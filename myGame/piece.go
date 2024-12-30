package myGame

const (
	// White is a constant that represents the color white
	White = "white"
	// Black is a constant that represents the color black
	Black = "black"
)

const (
	PAWN   = "pawn"
	KNIGHT = "knight"
	BISHOP = "bishop"
	ROOK   = "rook"
	QUEEN  = "queen"
	KING   = "king"
)

var icon = map[string]map[string]string{
	"black": {
		"pawn":   "♙",
		"knight": "♘",
		"bishop": "♗",
		"rook":   "♖",
		"queen":  "♕",
		"king":   "♔",
	},
	"white": {
		"pawn":   "♟",
		"knight": "♞",
		"bishop": "♝",
		"rook":   "♜",
		"queen":  "♛",
		"king":   "♚",
	},
}

type Piece struct {
	Color         string
	Type          string
	Icon          string
	Cell          *Cell
	PossibleMoves []Move
}

type Move struct {
	Piece Piece
	From  *Cell
	To    *Cell
}

func NewPiece(color, pieceType string) *Piece {
	p := Piece{Color: color, Type: pieceType}
	p.Icon = icon[color][pieceType]
	return &p
}
