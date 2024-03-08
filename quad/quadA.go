package main

import "github.com/01-edu/z01"

// QuadA
func QuadA(x, y int) {
	if x > 0 && y > 0 {
		// Loop through rows
		for i := 1; i <= y; i++ {
			// Check if it's the first or last row
			if i == 1 || i == y {
				// Print top or bottom row
				for j := 1; j <= x; j++ {
					if j == 1 || j == x {
						z01.PrintRune('o')
					} else {
						z01.PrintRune('-')
					}
				}
				z01.PrintRune('\n')
			} else {
				// Print middle rows
				for j := 1; j <= x; j++ {
					if j == 1 || j == x {
						z01.PrintRune('|')
					} else {
						z01.PrintRune(' ')
					}
				}
				z01.PrintRune('\n')
			}
		}
	}
}
func main() {
	QuadA(5, 3)
}
