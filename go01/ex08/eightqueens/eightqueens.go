package eightqueens

import "fmt"

func EightQueens() {
	solutions := make([][]int, 0)
	board := make([]int, 8)
	SolveQueens(board, 0, &solutions)
	printSolutions(solutions)
}

func SolveQueens(board []int, row int, solutions *[][]int) {
	if row == 8 {
		solution := make([]int, 8)
		copy(solution, board)
		*solutions = append(*solutions, solution)
		return
	}

	for col := 0; col < 8; col++ {
		if isValid(board, row, col) {
			board[row] = col
			SolveQueens(board, row+1, solutions)
		}
	}
}

func isValid(board []int, row, col int) bool {
	for i := 0; i < row; i++ {
		if board[i] == col ||
			board[i]-i == col-row ||
			board[i]+i == col+row {
			return false
		}
	}
	return true
}

func printSolutions(solutions [][]int) {
	for _, solution := range solutions {
		for _, col := range solution {
			fmt.Print(col + 1)
		}
		fmt.Println()
	}
}
