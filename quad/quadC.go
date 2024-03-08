package main
import "github.com/01-edu/z01"

// QuadC
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
func main() {
	QuadC(5,3)
}
