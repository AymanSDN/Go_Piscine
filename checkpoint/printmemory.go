package main

import (
	"unicode"

	"github.com/01-edu/z01"
)

func PrintStr(s string) {
	for _, char := range s {
		z01.PrintRune(char)
	}
}

func asciiToHex(s byte) string {
	var result string
	tmp := []byte{0, 0}
	tmp[0], tmp[1] = byte(s/16), byte(s%16)
	for i := 0; i < 2; i++ {
		if tmp[i] < 10 {
			result += string(tmp[i] + '0')
		} else {
			result += string(tmp[i] - 10 + 'a')
		}
	}
	return result
}



func PrintMemory(arr [10]byte) {
	var rst1 string
	var rst2 string
	size := len(arr)
	for i := 0; i < size; i++ {
		if unicode.IsGraphic(rune(arr[i])) {
			rst2 += string(arr[i])
		} else {
			rst2 += "."
		}
		rst1 += asciiToHex(arr[i])
		if i%4 == 3 && i != 0 {
			rst1 += "\n"
		} else {
			if i < size-1 {
				rst1 += " "
			}
		}
	}
	PrintStr(rst1+"\n"+rst2+"\n")
}

func main() {
	PrintMemory([10]byte{'h', 'e', 'l', 'l', 'o', 16, 21, '*'})
}
