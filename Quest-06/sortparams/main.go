package main

import (
	"os"

	"github.com/01-edu/z01"
)

func SortIntegerTable(table []int) {
	size := len(table)
	for i := 0; i < size; i++ {
		for j := 0; j < size-1-i; j++ {
			if table[j] > table[j+1] {
				tmp := table[j]
				table[j] = table[j+1]
				table[j+1] = tmp
			}
		}
	}
}

func main() {
	args := os.Args
	var array []int
	for i := 1; i < len(args); i++ {
		param := []rune(args[i])
		array = append(array, int(param[0]))
	}
	SortIntegerTable(array)
	for i := 0; i < len(array); i++ {
		z01.PrintRune(rune(array[i]))
		z01.PrintRune('\n')
	}
}
