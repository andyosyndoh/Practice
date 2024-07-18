package main

import (
	"fmt"
	"os"
	"os/exec"
)

const (
	EMPTY = " "
	X     = "X"
	O     = "O"
)

var board [3][3]string

var currentPlayer string

func main() {
	clearScreen()
	initializeBoard()
	currentPlayer = X
	for {
		clearScreen()
		printBoard()
		row, col := getPlayerMove()
		makeMove(row, col)
		if isGameOver(row, col) {
			printBoard()
			fmt.Printf("Player %s wins!\n", currentPlayer)
			os.Exit(0)
		}
		if checkDraw() {
			printBoard()
			fmt.Println("GameOver")
			os.Exit(0)
		}
		switchPlayer()
	}
}

func clearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func initializeBoard() {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			board[i][j] = EMPTY
		}
	}
}

func printBoard() {
	fmt.Println("  0 1 2")
	for i := 0; i < 3; i++ {
		fmt.Printf("%d ", i)
		for j := 0; j < 3; j++ {
			fmt.Printf("%s ", board[i][j])
		}
		fmt.Println()
	}
}

func getPlayerMove() (int, int) {
	var row, col int
	fmt.Printf("Player %s's turn. Enter row (0-2): ", currentPlayer)
	fmt.Scan(&row)
	if row < 0 || row > 2 {
		fmt.Printf("Player %s to repeat please input 0,1 0r 2: ", currentPlayer)
		fmt.Scan(&row)
		if row < 0 || row > 2 {
			fmt.Println("Game terminated because of breaking rules")
			os.Exit(1)
		}
	}
	fmt.Printf("Player %s's turn. Enter column (0-2): ", currentPlayer)
	fmt.Scan(&col)
	if col < 0 || col > 2 {
		fmt.Printf("Player %s to repeat please input 0,1 0r 2: ", currentPlayer)
		fmt.Scan(&col)
		if col < 0 || col > 2 {
			fmt.Println("Game terminated because of breaking rules")
			os.Exit(1)
		}
	}
	return row, col
}

func makeMove(row, col int) {
	if board[row][col] == EMPTY {
		board[row][col] = currentPlayer
	}
}

func switchPlayer() {
	if currentPlayer == X {
		currentPlayer = O
	} else {
		currentPlayer = X
	}
}

func isGameOver(row, col int) bool {
	return checkRow(row) || checkColumn(col) || checkDiagonals()
}

func checkRow(row int) bool {
	return board[row][0] == currentPlayer && board[row][1] == currentPlayer && board[row][2] == currentPlayer
}

func checkColumn(col int) bool {
	return board[0][col] == currentPlayer && board[1][col] == currentPlayer && board[2][col] == currentPlayer
}

func checkDiagonals() bool {
	return (board[0][0] == currentPlayer && board[1][1] == currentPlayer && board[2][2] == currentPlayer) ||
		(board[0][2] == currentPlayer && board[1][1] == currentPlayer && board[2][0] == currentPlayer)
}

func checkDraw() bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[i][j] == EMPTY {
				return false
			}
		}
	}
	return true
}
