package main

import (
	"os"
	"sudoku/solver"
	"sudoku/utils"
)


func main() {
	args := os.Args[1:]
	size := len(args)
	if size == 9 {
		Sudoku := utils.MakeCopy(args, size)
		if solver.SudokuSolver(Sudoku, 0, 0) {
			utils.PrintSudoku(Sudoku)
		} else {
			utils.PrintStr("Error\n")
		}
	} else {
		utils.PrintStr("Error\n")
	}
}
