package main

import "github.com/01-edu/z01"

func PrintComb2() {
	for a := '0'; a <= '9'; a++ {
		for b := '0'; b <= '9'; b++ {
			for c := '0'; c <= '9'; c++ {
				for d := '1'; d <= '9'; d++ {
					if a < c || (b < d && a == c) {
						z01.PrintRune(a)
						z01.PrintRune(b)
						z01.PrintRune(' ')
						z01.PrintRune(c)
						z01.PrintRune(d)
						if !(a == '9' && b == '8') {
							z01.PrintRune(',')
							z01.PrintRune(' ')
						}

					}
				}
			}
		}
	}
	z01.PrintRune('\n')
}

// func main() {
// 	PrintComb2()
// }
