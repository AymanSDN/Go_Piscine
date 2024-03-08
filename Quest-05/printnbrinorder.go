package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	if n == 0 {
		z01.PrintRune('0')
	}
	var tab [10]int
	for n > 0 {
		index := n % 10
		tab[index] += 1
		n /= 10
	}
	for i := 0; i < 10; i++ {
		for tab[i] > 0 {
			z01.PrintRune(rune(i + 48))
			tab[i]--
		}
	}
}
