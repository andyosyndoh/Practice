package main

import "fmt"

func main() {
	board := createBoard()
	if PlaceRecursive(board, 0) {
		printBoard(board)
	}
}

func createBoard() [][]string {
	board := make([][]string, 8)
	for i := range board {
		board[i] = make([]string, 8)
		for j := range board[i] {
			board[i][j] = "."
		}
	}
	return board
}

func printBoard(board [][]string) {
	for i := range board {
		for j := range board[i] {
			fmt.Print(board[i][j])
			if j < len(board[i])-1 {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func PlaceRecursive(board [][]string, queen int) bool {
	if queen == 8 {
		return true
	}
	for i := range board {
		for j := range board[i] {
			if canplace(board, i, j) {
				Place(board, i, j)
				if PlaceRecursive(board, queen+1) {
                    return true
                }
				Remove(board, i, j)
			}
		}
	}
	return false
}

func canplace(board [][]string, row, col int) bool {
	for i := 0; i < 7; i++ {
		if board[row][i] == "Q" {
            return false
        }
	}
	for i := 0; i < 7; i++ {
		if board[i][col] == "Q" {
            return false
        }
	}
	for i, j := row, col; i >= 0 && j >= 0; i, j = i-1, j-1 {
        if board[i][j] == "Q" {
            return false
        }
    }
	for i, j := row, col; i >= 0 && j < 8; i, j = i-1, j+1 {
        if board[i][j] == "Q" {
            return false
        }
    }
	for i, j := row, col; i < 8 && j >= 0; i, j = i+1, j-1 {
        if board[i][j] == "Q" {
            return false
        }
    }
	for i, j := row, col; i < 8 && j < 8; i, j = i+1, j+1 {
        if board[i][j] == "Q" {
            return false
        }
    }
	return true

}

func Place(board [][]string, row, col int) {
    board[row][col] = "Q"
}

func Remove(board [][]string, row, col int) {
    board[row][col] = "."
}


