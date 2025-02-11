package core

import "fmt"

type Game struct {
	Board   *Board
	Players []Player
	Current int
}

func NewGame(size int, players []Player) *Game {
	return &Game{
		Board:   NewBoard(size),
		Players: players,
		Current: 0,
	}
}

func (g *Game) Play() {
	for !g.Board.IsFull() {
		g.Board.Display()
		currentPlayer := g.Players[g.Current]

		fmt.Printf("Player %s, enter your move (row col): ", currentPlayer.Name)
		row, col := GetUserInput()

		if !g.Board.PlaceSymbol(row, col, currentPlayer.Symbol) {
			fmt.Println("Invalid move, try again!")
			continue
		}

		if g.Board.CheckWinner(currentPlayer.Symbol) {
			g.Board.Display()
			fmt.Printf("Player %s wins! 🎉\n", currentPlayer.Name)
			return
		}

		g.Current = (g.Current + 1) % len(g.Players)
	}
	fmt.Println("It's a draw! 🤝")
}
