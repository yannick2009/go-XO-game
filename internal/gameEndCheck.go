package internal

import "github.com/go-terminal-OX/pkg/vars"

// CheckWinner is a function that checks if there is a winner in the game
func CheckWinner() bool {
	return checkRows() || checkColumns() || checkDiagonals()
}

// CheckGameOver is a function that checks if the game is over
func CheckGameOver() bool {
	return vars.RemainingMoves == 0
}

// CheckDiagonals is a function that checks the diagonals of the game board for a winner
func checkDiagonals() bool {
	board := vars.GameBoard
	// Check main diagonal (top-left to bottom-right)
	if board[0][0] == board[1][1] && board[1][1] == board[2][2] && board[2][2] != "" {
		return true
	}
	// Check secondary diagonal (top-right to bottom-left)
	if board[0][2] == board[1][1] && board[1][1] == board[2][0] && board[2][0] != "" {
		return true
	}
	return false
}

// CheckRows is a function that checks the rows of the game board for a winner
func checkRows() bool {
	board := vars.GameBoard
	// Check each row
	for row := 0; row < 3; row++ {
		if board[row][0] == board[row][1] && board[row][1] == board[row][2] && board[row][2] != "" {
			return true
		}
	}
	return false
}

// CheckColumns is a function that checks the columns of the game board for a winner
func checkColumns() bool {
	board := vars.GameBoard
	// Check each column
	for col := 0; col < 3; col++ {
		if board[0][col] == board[1][col] && board[1][col] == board[2][col] && board[2][col] != "" {
			return true
		}
	}
	return false
}
