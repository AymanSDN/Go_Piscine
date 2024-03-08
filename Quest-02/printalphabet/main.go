package main

import "github.com/01-edu/z01"

// this program print from 'a' to 'z'

func main() {
	for char := 'a'; char <= 'z'; char++ {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}