package main

import (
	"os"

	"github.com/01-edu/z01"
)

func PrintStr(s string) {
	for _, char := range s {
		z01.PrintRune(char)
	}
}

func removeDuplicates(str string) string {
    var result string

    for _, char := range str {
        if !containsChar(result, char) {
            result += string(char)
        }
    }

    return result
}

func containsChar(str string, char rune) bool {
    for _, c := range str {
        if c == char {
            return true
        }
    }
    return false
}

func main() {
	if len(os.Args) == 3 {
		s := removeDuplicates(removeDuplicates(os.Args[1]) + removeDuplicates(os.Args[2]))
		PrintStr(s)
	}
	z01.PrintRune('\n')

}
