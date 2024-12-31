package myGame

func (b *Board) PossibleMoves(p *Piece) []Move {
	switch p.Type {
	case PAWN:
		return b.pawnMoves(p)
	case KNIGHT:
		return b.knightMoves(p)
	case BISHOP:
		return b.bishopMoves(p)
	case ROOK:
		return b.rookMoves(p)
	case QUEEN:
		return b.queenMoves(p)
	case KING:
		return b.kingMoves(p)
	}
	return []Move{}
}

func inBoard(row, column int) bool {
	return row >= 0 && row < 8 && column >= 0 && column < 8
}

func canMoveTo(cell *Cell, color string) bool {
	return cell.Piece == nil || cell.Piece.Color != color
}

func (b *Board) pawnMoves(p *Piece) []Move {
	moves := []Move{}
	curr := p.Cell
	color := p.Color
	move := 1
	// if white, move up
	dx := 1
	// if black, move down
	if color == Black {
		dx = -1
	}
	// if first move, can move 2 spaces
	if (color == White && curr.Row == 1) || (color == Black && curr.Row == 6) {
		move = 2
	}
	for i := 1; i <= move; i++ {
		newCell := b.Cells[curr.Row+dx*i][curr.Column]
		if newCell.Piece != nil {
			break
		}
		moves = append(moves, Move{Piece: *p, From: curr, To: &newCell})
	}
	// if enemy piece is diagonal, can move diagonally
	if curr.Column > 0 {
		newCell := b.Cells[curr.Row+dx][curr.Column-1]
		if newCell.Piece != nil && newCell.Piece.Color != color {
			moves = append(moves, Move{Piece: *p, From: curr, To: &newCell})
		}
	}
	if curr.Column < 7 {
		newCell := b.Cells[curr.Row+dx][curr.Column+1]
		if newCell.Piece != nil && newCell.Piece.Color != color {
			moves = append(moves, Move{Piece: *p, From: curr, To: &newCell})
		}
	}
	// if enemy piece is in front, can move diagonally
	return moves
}

func (b *Board) knightMoves(p *Piece) []Move {
	moves := []Move{}
	possibleMoves := [][]int{
		{-2, -1},
		{-2, 1},
		{-1, -2},
		{-1, 2},
		{1, -2},
		{1, 2},
		{2, -1},
		{2, 1},
	}
	for _, move := range possibleMoves {
		newRow := p.Cell.Row + move[0]
		newColumn := p.Cell.Column + move[1]
		if inBoard(newRow, newColumn) && canMoveTo(&b.Cells[newRow][newColumn], p.Color) {
			moves = append(moves, Move{Piece: *p, From: p.Cell, To: &b.Cells[newRow][newColumn]})
		}
	}
	// move in L shape
	// can jump over pieces
	return moves
}

func (b *Board) bishopMoves(p *Piece) []Move {
	moves := []Move{}
	// move diagonally
	// can't jump over pieces
	for _, move := range [][]int{{-1, -1}, {-1, 1}, {1, -1}, {1, 1}} {
		for {
			newRow := p.Cell.Row + move[0]
			newColumn := p.Cell.Column + move[1]
			if !inBoard(newRow, newColumn) {
				break
			}
			newCell := b.Cells[newRow][newColumn]
			if newCell.Piece != nil {
				if newCell.Piece.Color != p.Color {
					moves = append(moves, Move{Piece: *p, From: p.Cell, To: &newCell})
				}
				break
			}
			moves = append(moves, Move{Piece: *p, From: p.Cell, To: &newCell})
		}
	}
	return moves
}

func (b *Board) rookMoves(p *Piece) []Move {
	moves := []Move{}
	for _, move := range [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
		for {
			newRow := p.Cell.Row + move[0]
			newColumn := p.Cell.Column + move[1]
			if !inBoard(newRow, newColumn) {
				break
			}
			newCell := b.Cells[newRow][newColumn]
			if newCell.Piece != nil {
				if newCell.Piece.Color != p.Color {
					moves = append(moves, Move{Piece: *p, From: p.Cell, To: &newCell})
				}
				break
			}
			moves = append(moves, Move{Piece: *p, From: p.Cell, To: &newCell})
		}
	}
	// move vertically or horizontally
	// can't jump over pieces
	return moves
}

func (b *Board) queenMoves(p *Piece) []Move {
	moves := []Move{}
	// can move like a bishop or a rook
	moves = append(moves, b.bishopMoves(p)...)
	moves = append(moves, b.rookMoves(p)...)
	// move vertically, horizontally, or diagonally
	// can't jump over pieces
	return moves
}

func (b *Board) kingMoves(p *Piece) []Move {
	moves := []Move{}
	possibleMoves := [][]int{
		{-1, -1},
		{-1, 0},
		{-1, 1},
		{0, -1},
		{0, 1},
		{1, -1},
		{1, 0},
		{1, 1},
	}
	for _, move := range possibleMoves {
		newRow := p.Cell.Row + move[0]
		newColumn := p.Cell.Column + move[1]
		if inBoard(newRow, newColumn) && canMoveTo(&b.Cells[newRow][newColumn], p.Color) {
			moves = append(moves, Move{Piece: *p, From: p.Cell, To: &b.Cells[newRow][newColumn]})
		}
	}
	// move one space in any direction
	// can't move into check
	return moves
}
