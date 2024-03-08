package main

import (
	"os"
	"quadchecker/utils"

	"github.com/01-edu/z01"
)

// QuadA prints a rectangle with corners marked as 'o' and edges marked as '-'
func QuadA(x, y int) {
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

// QuadB prints a rectangle with diagonal lines from top left to bottom right and from top right to bottom left
func QuadB(x, y int) {
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

// QuadC prints a rectangle with letters 'A', 'B', and 'C' marking the corners and edges
func QuadC(x, y int) {
	// Check if dimensions are valid
	if x > 0 && y > 0 {
		// Loop through rows
		for i := 1; i <= y; i++ {
			// Loop through columns
			for j := 1; j <= x; j++ {
				// Determine character to print based on position
				if (i == 1 && j == 1) || (i == 1 && j == x) {
					z01.PrintRune('A') // Top-left and top-right corners
				} else if (i == y && j == x) || (i == y && j == 1) {
					z01.PrintRune('C') // Bottom-right and bottom-left corners
				} else if (j == 1 || j == x) || (i == 1 || i == y) {
					z01.PrintRune('B') // Vertical edges and horizontal edges (excluding corners)
				} else {
					z01.PrintRune(' ') // Inner space
				}
			}
			z01.PrintRune('\n') // Move to the next line after each row
		}
	}
}

// QuadD prints a rectangle with letters 'A', 'B', and 'C' marking the corners and edges
func QuadD(x, y int) {
	// Check if dimensions are valid
	if x > 0 && y > 0 {
		// Loop through rows
		for i := 1; i <= y; i++ {
			// Loop through columns
			for j := 1; j <= x; j++ {
				// Determine character to print based on position
				if (i == 1 && j == 1) || (i == y && j == 1) {
					z01.PrintRune('A') // Top-left and bottom-left corners
				} else if (i == y && j == x) || (i == 1 && j == x) {
					z01.PrintRune('C') // Bottom-right and top-right corners
				} else if (j == 1 || j == x) || (i == 1 || i == y) {
					z01.PrintRune('B') // Vertical edges and horizontal edges (excluding corners)
				} else {
					z01.PrintRune(' ') // Inner space
				}
			}
			z01.PrintRune('\n') // Move to the next line after each row
		}
	}
}

// QuadE prints a rectangle with letters 'A', 'B', and 'C' marking the corners and edges
func QuadE(x, y int) {
	// Check if dimensions are valid
	if x > 0 && y > 0 {
		// Loop through rows
		for i := 1; i <= y; i++ {
			// Loop through columns
			for j := 1; j <= x; j++ {
				// Determine character to print based on position
				if (i == 1 && j == 1) || (i == y && j == x && i != 1 && j != 1) {
					z01.PrintRune('A') // Top-left corner and bottom-right corner (except for first row)
				} else if (i == 1 && j == x) || (i == y && j == 1) {
					z01.PrintRune('C') // Top-right corner and bottom-left corner
				} else if (j == 1 || j == x) || (i == 1 || i == y) {
					z01.PrintRune('B') // Vertical edges and horizontal edges (excluding corners)
				} else {
					z01.PrintRune(' ') // Inner space
				}
			}
			z01.PrintRune('\n') // Move to the next line after each row
		}
	}
}

func main() {

	if len(os.Args) == 3 {
		x := utils.Atoi(os.Args[1])
		y := utils.Atoi(os.Args[2])
		QuadA(x, y)
		// QuadB(x, y)
		// QuadC(x, y)
		// QuadD(x, y)
		// QuadE(x, y)
	}
}
