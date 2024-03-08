package main

import "github.com/01-edu/z01"

// This function check if nb passed as param is negative

func IsNegative(nb int) {
	if nb < 0 {
		z01.PrintRune('T')
	} else {
		z01.PrintRune('F')
	}
}

// func main() {
// 	IsNegative(5)
// }
