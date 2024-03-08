package utils

import "github.com/01-edu/z01"

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
func PrintStr(s string) {
	for _, char := range s {
		z01.PrintRune(char)
	}
}
func GridtoString(grid []string) string {
	var result []rune
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			result = append(result, rune(grid[i][j]))
		}
	}
	return string(result)
}

// QuadA prints a rectangle with corners marked as 'o' and edges marked as '-'
func QuadA(x, y int) string {
	var rst []rune
	// Loop through rows
	for i := 1; i <= y; i++ {
		// Check if it's the first or last row
		if i == 1 || i == y {
			// Print top or bottom row
			for j := 1; j <= x; j++ {
				if j == 1 || j == x {
					rst = append(rst, 'o')
				} else {
					rst = append(rst, '-')
				}
			}
		} else {
			// Print middle rows
			for j := 1; j <= x; j++ {
				if j == 1 || j == x {
					rst = append(rst, '|')
				} else {
					rst = append(rst, ' ')
				}
			}
		}
	}
	return string(rst)
}

// QuadB prints a rectangle with diagonal lines from top left to bottom right and from top right to bottom left
func QuadB(x, y int) string {
	var rst []rune
	// Loop through rows
	for i := 1; i <= y; i++ {
		// Check if it's the first row
		if i == 1 {
			// Print top row with top-left and top-right corners
			for j := 1; j <= x; j++ {
				if j == 1 {
					rst = append(rst, '/')
				} else if j == x {
					rst = append(rst, '\\')
				} else { // z01.PrintRune('o')
					rst = append(rst, '*')
				}
			}
		} else if i == y {
			// Print bottom row with bottom-left and bottom-right corners
			for j := 1; j <= x; j++ {
				if j == 1 {
					rst = append(rst, '\\')
				} else if j == x {
					rst = append(rst, '/')
				} else {
					rst = append(rst, '*')
				}
			}
		} else {
			// Print middle rows with '*' at edges
			for j := 1; j <= x; j++ {
				if j == 1 || j == x {
					rst = append(rst, '*')
				} else {
					rst = append(rst, ' ')
				}
			}
		}
	}
	return string(rst)
}

// QuadC prints a rectangle with letters 'A', 'B', and 'C' marking the corners and edges
func QuadC(x, y int) string {
	var rst []rune
	// Check if dimensions are valid
	if x > 0 && y > 0 {
		// Loop through rows
		for i := 1; i <= y; i++ {
			// Loop through columns
			for j := 1; j <= x; j++ {
				// Determine character to print based on position
				if (i == 1 && j == 1) || (i == 1 && j == x) {
					rst = append(rst, 'A') // Top-left and top-right corners
				} else if (i == y && j == x) || (i == y && j == 1) {
					rst = append(rst, 'C') // Bottom-right and bottom-left corners
				} else if (j == 1 || j == x) || (i == 1 || i == y) {
					rst = append(rst, 'B') // Vertical edges and horizontal edges (excluding corners)
				} else {
					rst = append(rst, ' ') // Inner space
				}
			}
		}
	}
	return string(rst)
}

// QuadD prints a rectangle with letters 'A', 'B', and 'C' marking the corners and edges
func QuadD(x, y int) string {
	var rst []rune
	// Check if dimensions are valid
	if x > 0 && y > 0 {
		// Loop through rows
		for i := 1; i <= y; i++ {
			// Loop through columns
			for j := 1; j <= x; j++ {
				// Determine character to print based on position
				if (i == 1 && j == 1) || (i == y && j == 1) {
					rst = append(rst, 'A') // Top-left and bottom-left corners
				} else if (i == y && j == x) || (i == 1 && j == x) {
					rst = append(rst, 'C') // Bottom-right and top-right corners
				} else if (j == 1 || j == x) || (i == 1 || i == y) {
					rst = append(rst, 'B') // Vertical edges and horizontal edges (excluding corners)
				} else {
					rst = append(rst, ' ') // Inner space
				}
			}
		}
	}
	return string(rst)
}

// QuadE prints a rectangle with letters 'A', 'B', and 'C' marking the corners and edges
func QuadE(x, y int) string {
	var rst []rune
	// Check if dimensions are valid
	if x > 0 && y > 0 {
		// Loop through rows
		for i := 1; i <= y; i++ {
			// Loop through columns
			for j := 1; j <= x; j++ {
				// Determine character to print based on position
				if (i == 1 && j == 1) || (i == y && j == x && i != 1 && j != 1) {
					rst = append(rst, 'A') // Top-left corner and bottom-right corner (except for first row)
				} else if (i == 1 && j == x) || (i == y && j == 1) {
					rst = append(rst, 'C') // Top-right corner and bottom-left corner
				} else if (j == 1 || j == x) || (i == 1 || i == y) {
					rst = append(rst, 'B') // Vertical edges and horizontal edges (excluding corners)
				} else {
					rst = append(rst, ' ') // Inner space
				}
			}
		}
	}
	return string(rst)
}
