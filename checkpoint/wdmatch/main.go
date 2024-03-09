package main

import (
	// "fmt"
	"os"

	"github.com/01-edu/z01"
)

func PrintStr(s string) {
	for _, char := range s {
		z01.PrintRune(char)
	}
}

func wdmatch(tomatch string, str string) bool {
	tm_size := len(tomatch)
	size := len(str)
	flag := 0
	k := 0
	for i := 0; i < tm_size; i++ {
		for j := k; j < size; j++ {
			if tomatch[i] == str[j] {
				flag++
				k = j
				break
			}
		}
	}
	return flag == tm_size
}

func main() {
	if len(os.Args) == 3 {
		if wdmatch(os.Args[1], os.Args[2]) {
			PrintStr(os.Args[1])
		}
	}
	PrintStr("\n")
}
