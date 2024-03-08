package main

import "github.com/01-edu/z01"

// QuadB
func QuadB(x, y int) {
	if x > 0 && y > 0 {
		// Loop through rows
		for i := 1; i <= y; i++ {
			// Check if it's the first row
			if i == 1 {
				// Print top row with top-left and top-right corners
				for j := 1; j <= x; j++ {
					if j == 1 {
						z01.PrintRune('/')
					} else if j == x {
						z01.PrintRune('\\')
					} else {
						z01.PrintRune('*')
					}
				}
			} else if i == y {
				// Print bottom row with bottom-left and bottom-right corners
				for j := 1; j <= x; j++ {
					if j == 1 {
						z01.PrintRune('\\')
					} else if j == x {
						z01.PrintRune('/')
					} else {
						z01.PrintRune('*')
					}
				}
			} else {
				// Print middle rows with '*' at edges
				for j := 1; j <= x; j++ {
					if j == 1 || j == x {
						z01.PrintRune('*')
					} else {
						z01.PrintRune(' ')
					}
				}
			}
			z01.PrintRune('\n')
		}
	}
}
func main() {
	QuadB(5, 3)
}
