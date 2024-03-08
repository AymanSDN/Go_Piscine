package main

import "github.com/01-edu/z01"

// this program print from '0' to '9'

func main() {
	for char := '0'; char <= '9'; char++ {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}
