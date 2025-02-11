package core

import "fmt"

type Board struct {
	Size  int
	Cells [][]PlayerSymbol
}

func NewBoard(size int) *Board {
	board := &Board{
		Size:  size,
		Cells: make([][]PlayerSymbol, size),
	}
	for i := range board.Cells {
		board.Cells[i] = make([]PlayerSymbol, size)
	}
	return board
}

func (b *Board) Display() {
	fmt.Println("\nCurrent Board:")
	for _, row := range b.Cells {
		for _, cell := range row {
			if cell != nil {
				fmt.Printf(" %s ", cell.GetSymbol())
			} else {
				fmt.Print(" - ")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}
