package myGame

const (
	// columns a to h
	a = 0
	b = 1
	c = 2
	d = 3
	e = 4
	f = 5
	g = 6
	h = 7
)

type Cell struct {
	Row    int
	Column int
	Piece  *Piece
}
