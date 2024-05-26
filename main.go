package main

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"time"
)

func clearScreen() {
	fmt.Print("\033[2J")
	fmt.Print("\033[H")
}

func main() {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)

	go func() {
		<-signalChan
		fmt.Println("\nQuitting Game")
		os.Exit(0)
	}()

	state := NewGameState()

	state.InitializeGame()

	reader := bufio.NewReader(os.Stdin)

	for {
		clearScreen()

		fmt.Printf("2048 - Score: %d\n", state.score)
		state.PrintBoard()

		fmt.Print("Move (WASD to shift the board, Q to quit): ")
		command, _ := reader.ReadString('\n')

		switch command[0] {
		case 'w':
			state.ShiftUp()
		case 's':
			state.ShiftDown()
		case 'a':
			state.ShiftLeft()
		case 'd':
			state.ShiftRight()
		case 'q':
			os.Exit(0)
		default:
			fmt.Println("Invalid command. Command must be W, A, S, D, or Q")
			time.Sleep(1 * time.Second)
			continue
		}

		if state.Has2048() {
			clearScreen()
			fmt.Printf("You Won!. Score: %d\n", state.score)
			state.PrintBoard()
			break
		} else if !state.IsValid() {
			clearScreen()
			fmt.Printf("You lose!. There are no valid moves left\n")
			state.PrintBoard()
			break
		}

		state.PlaceRandomTile()
	}

}
