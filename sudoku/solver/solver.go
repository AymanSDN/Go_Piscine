package solver


// CanIPut checks if it's possible to put a value in a specific position of the Sudoku grid
func CanIPut(sudoku [][]int, col int, row int, value int) bool {
    return !CheckVertical(sudoku, col, row, value) && !CheckHorizontal(sudoku, col, row, value) && !CheckSquare3_3(sudoku, col, row, value)
}

// CheckVertical checks if the same value exists in the same column of the current position
func CheckVertical(sudoku [][]int, col int, row int, value int) bool {
    for i := 0; i < 9; i++ {
        if sudoku[i][col] == value {
            return true
        }
    }
    return false
}

// CheckHorizontal checks if the same value exists in the same row of the current position
func CheckHorizontal(sudoku [][]int, col int, row int, value int) bool {
    for i := 0; i < 9; i++ {
        if sudoku[row][i] == value {
            return true
        }
    }
    return false
}

// CheckSquare3_3 checks if placing a value in the 3x3 square containing the cell at (col, row) in the Sudoku puzzle is valid.
// It returns true if the value conflicts with any existing value in the square, and false otherwise.
func CheckSquare3_3(sudoku [][]int, col int, row int, value int) bool {
    // Calculate the top-left corner coordinates of the 3x3 square containing the cell at (col, row).
    scol, srow := int(col/3)*3, int(row/3)*3
    // Iterate through each cell in the 3x3 square.
    for drow := 0; drow < 3; drow++ {
        for dcol := 0; dcol < 3; dcol++ {
            // If the value already exists in any cell of the square, return true (indicating a conflict).
            if sudoku[srow+drow][scol+dcol] == value {
                return true
            }
        }
    }
    // If the value doesn't conflict with any existing value in the square, return false.
    return false
}


// NextEmptyCell returns the coordinates of the next empty cell in the Sudoku puzzle,
func NextEmptyCell(col int, row int) (int, int) {
    // Calculate the column index of the next cell.
    nextCol := (col + 1) % 9
    // If the next column index becomes 0, it means we've reached the end of the row,
    // so move to the next row.
    nextRow := row
    if nextCol == 0 {
        nextRow = row + 1
    }
    // Return the coordinates of the next empty cell.
    return nextCol, nextRow
}


// SudokuSolver is a recursive function to solve a Sudoku puzzle.
// It uses backtracking algorithm to try different values for each empty cell.
func SudokuSolver(sudoku [][]int, col int, row int) bool {
    // If we've reached the end of the puzzle (row == 9), it means we have successfully solved it.
    if row == 9 {
        return true
    }
    // If the current cell is not empty, move to the next empty cell.
    if sudoku[row][col] != 0 {
        tmpx, tmpy := NextEmptyCell(col, row)
        return SudokuSolver(sudoku, tmpx, tmpy)
    } else {
        // If the current cell is empty, try filling it with values from 1 to 9.
        for i := 0; i < 9; i++ {
            value := i + 1
            // Check if the current value can be placed in the current cell.
            if CanIPut(sudoku, col, row, value) {
                sudoku[row][col] = value
                // Move to the next empty cell recursively.
                tmpx, tmpy := NextEmptyCell(col, row)
                if SudokuSolver(sudoku, tmpx, tmpy) {
                    return true
                }
                // If placing the current value doesn't lead to a solution, backtrack and try the next value.
                sudoku[row][col] = 0
            }
        }
        // If none of the values lead to a solution, backtrack.
        return false
    }
}

