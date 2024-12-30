package myGame

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func StartGame() {
	b := NewBoard()
	b.PrintBoard(false)

	for !b.GameOver() {
		// ask for move as input in terminal in form of "e2 e4"
		curr, new := getMove(b)
		// if valid, make the move
		error := b.MakeMove(curr, new)
		if error != nil {
			continue
		}
		b.PrintBoard(b.CurrentPlayer != White)
	}
}

func getMove(b *Board) (string, string) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter curr pos for move for player %s (e.g. 'e2'): ", b.CurrentPlayer)

	curPos, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return getMove(b)
	}

	// Clean the input
	curPos = strings.TrimSpace(curPos)

	fmt.Printf("Enter new pos for move for player %s (e.g. 'e2'): ", b.CurrentPlayer)

	newPos, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return getMove(b)
	}

	// Clean the input
	newPos = strings.TrimSpace(newPos)

	return curPos, newPos
}
