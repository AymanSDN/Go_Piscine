package main

import (
	"bufio"
	"fmt"
	"os"
	"quadchecker/utils"
	"strconv"
)

// getProgramOutput reads lines from standard input and returns them as a slice of strings.
// It also returns the dimensions of the grid and a boolean indicating if the grid is not empty.
func getProgramOutput() ([]string, int, int, bool) {
	grid := []string{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, line)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprint(os.Stderr, "reading standard input:", err)
	}
	if len(grid) == 0 {
		return grid, 0, 0, false
	}
	x, y := len(grid[0]), len(grid)
	return grid, x, y, true
}

// QuadChecker checks if the given line matches any of the Quad functions (A to E).
// If a match is found, it appends the function name and its bounds to the result slice.
// If no match is found, it prints "Not a quad function".
func QuadChecker(line string, x, y int) {
	var result []string
	t_x := strconv.Itoa(x)
	t_y := strconv.Itoa(y)
	bound := "[" + t_x + "]" + "[" + t_y + "]"
	flag := 0

	if line == utils.QuadA(x, y) {
		result = append(result, "[QuadA]"+bound)
		flag++
	}
	if line == utils.QuadB(x, y) {
		result = append(result, "[QuadB]"+bound)
		flag++
	}
	if line == utils.QuadC(x, y) {
		result = append(result, "[QuadC]"+bound)
		flag++
	}
	if line == utils.QuadD(x, y) {
		result = append(result, "[QuadD]"+bound)
		flag++
	}
	if line == utils.QuadE(x, y) {
		result = append(result, "[QuadE]"+bound)
		flag++
	}
	if flag == 0 {
		utils.PrintStr("Not a quad function")
	}
	result = append(result, "\n")
	for j := 0; j < len(result); j++ {
		utils.PrintStr(result[j])
		if j < flag-1 {
			utils.PrintStr(" || ")
		}
	}
}

// main is the entry point of the program. It gets the program output and checks if it matches any Quad function.
func main() {
	grid, x, y, fl := getProgramOutput()
	if fl {
		line := utils.GridtoString(grid)
		QuadChecker(line, x, y)
	}
}
