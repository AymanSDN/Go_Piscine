package utils

import (

	"github.com/01-edu/z01"
)

func Atoi(s string) int {
	i := 0
	size := len(s)
	if size == 0 {
		return 0
	}
	isNegative := 1
	result := 0
	if s[i] == '-' || s[i] == '+' {
		if s[i] == '-' {
			isNegative = -1
		}
		i++
	}
	for i < size {
		if s[i] >= '0' && s[i] <= '9' {
			tmp := int(s[i] - '0')
			result = result*10 + tmp
		} else {
			return 0
		}
		i++
	}
	return result * isNegative
}

func PrintNbr(n int) {
	if n < 0 {
		z01.PrintRune('-')
		PrintNbr(-n)
	} else if n > 9 {
		PrintNbr(n / 10)
		PrintNbr(n % 10)
	} else {
		z01.PrintRune(rune(n + 48))
	}
}


func MakeCopy(args []string, size int) [][]int {
	Sudoku := make([][]int, size)
	for i := 0; i < size; i++ {
		Sudoku[i] = make([]int, len(args[i]))
		for j := 0; j < len(args[i]); j++ {
			Sudoku[i][j] = Atoi(string(args[i][j]))
		}
	}
	return Sudoku
}

func PrintStr(s string) {
	for i := 0; i < len(s); i++ {
		z01.PrintRune(rune(s[i]))
	}
}

func PrintSudoku(Sudoku [][]int) {
	for i := 0; i < len(Sudoku); i++ {
		for j := 0; j < len(Sudoku[i]); j++ {
			PrintNbr(Sudoku[i][j])
			if j < len(Sudoku[i])-1 {
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune('\n')
	}
}
