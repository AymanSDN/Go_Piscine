package main

import "github.com/01-edu/z01"

// this program print from 'z' to 'a'

func main() {
	for char := 'z'; char >= 'a'; char-- {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}
