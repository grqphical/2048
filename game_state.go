package main

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/fatih/color"
)

type GameState struct {
	board []int
	score int
}

func NewGameState() GameState {
	return GameState{
		board: make([]int, 16),
		score: 0,
	}
}

func (g *GameState) PlaceRandomTile() {
	for {
		pos := rand.Intn(16)

		if g.board[pos] == 0 {
			g.board[pos] = 2
			break
		}
	}
}

func (g *GameState) InitializeGame() {
	g.PlaceRandomTile()
	g.PlaceRandomTile()
}

func (g *GameState) ShiftRight() {
	for row := 0; row < 4; row++ {
		for col := 3; col >= 0; col-- {
			index := row*4 + col
			if g.board[index] != 0 {
				targetCol := col
				for targetCol < 3 && g.board[row*4+targetCol+1] == 0 {
					targetCol++
				}
				if targetCol != col {
					g.board[row*4+targetCol] = g.board[index]
					g.board[index] = 0
				}
				if targetCol < 3 && g.board[row*4+targetCol+1] == g.board[row*4+targetCol] {
					g.board[row*4+targetCol+1] *= 2
					g.board[row*4+targetCol] = 0

					g.score += g.board[row*4+targetCol+1]
				}
			}
		}
	}
}

func (g *GameState) ShiftLeft() {
	for row := 0; row < 4; row++ {
		for col := 0; col < 4; col++ {
			index := row*4 + col
			if g.board[index] != 0 {
				targetCol := col
				for targetCol > 0 && g.board[row*4+targetCol-1] == 0 {
					targetCol--
				}
				if targetCol != col {
					g.board[row*4+targetCol] = g.board[index]
					g.board[index] = 0
				}
				if targetCol > 0 && g.board[row*4+targetCol-1] == g.board[row*4+targetCol] {
					g.board[row*4+targetCol-1] *= 2
					g.board[row*4+targetCol] = 0

					g.score += g.board[row*4+targetCol-1]
				}
			}
		}
	}
}

func (g *GameState) ShiftUp() {
	for col := 0; col < 4; col++ {
		for row := 0; row < 4; row++ {
			index := row*4 + col
			if g.board[index] != 0 {
				targetRow := row
				for targetRow > 0 && g.board[(targetRow-1)*4+col] == 0 {
					targetRow--
				}
				if targetRow != row {
					g.board[targetRow*4+col] = g.board[index]
					g.board[index] = 0
				}
				if targetRow > 0 && g.board[(targetRow-1)*4+col] == g.board[targetRow*4+col] {
					g.board[(targetRow-1)*4+col] *= 2
					g.board[targetRow*4+col] = 0

					g.score += g.board[(targetRow-1)*4+col]
				}
			}
		}
	}
}

func (g *GameState) ShiftDown() {
	for col := 0; col < 4; col++ {
		for row := 3; row >= 0; row-- {
			index := row*4 + col
			if g.board[index] != 0 {
				targetRow := row
				for targetRow < 3 && g.board[(targetRow+1)*4+col] == 0 {
					targetRow++
				}
				if targetRow != row {
					g.board[targetRow*4+col] = g.board[index]
					g.board[index] = 0
				}
				if targetRow < 3 && g.board[(targetRow+1)*4+col] == g.board[targetRow*4+col] {
					g.board[(targetRow+1)*4+col] *= 2
					g.board[targetRow*4+col] = 0

					g.score += g.board[(targetRow+1)*4+col]
				}
			}
		}
	}
}

func (g *GameState) Has2048() bool {
	for _, v := range g.board {
		if v == 2048 {
			return true
		}
	}
	return false
}

func (g *GameState) IsValid() bool {
	// Check for any zeros
	for i := 0; i < 16; i++ {
		if g.board[i] == 0 {
			return true
		}
	}

	// Check for adjacent equal numbers horizontally
	for row := 0; row < 4; row++ {
		for col := 0; col < 3; col++ {
			if g.board[row*4+col] == g.board[row*4+col+1] {
				return true
			}
		}
	}

	// Check for adjacent equal numbers vertically
	for col := 0; col < 4; col++ {
		for row := 0; row < 3; row++ {
			if g.board[row*4+col] == g.board[(row+1)*4+col] {
				return true
			}
		}
	}

	return false
}

func (g *GameState) PrintBoard() {
	separator := "|------|------|------|------|"
	cyan := color.New(color.FgCyan).SprintFunc()
	fmt.Println(separator)
	for row := 0; row < 4; row++ {
		fmt.Print("|")
		for col := 0; col < 4; col++ {
			value := fmt.Sprintf("%d", g.board[row*4+col])
			// Pad the value to fit in 6 characters width
			if value != "0" {
				fmt.Printf(" %4s |", cyan(strings.Repeat(" ", 4-len(value))+value))
			} else {
				fmt.Printf(" %4s |", strings.Repeat(" ", 4-len(value))+value)
			}
		}
		fmt.Println()
		fmt.Println(separator)
	}
}
