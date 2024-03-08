package piscine

func NRune(s string, n int) rune {
	size := len(s)
	slice := []rune(s)
	if n <= 0 || size == 0 {
		return 0
	}
	if size >= n {
		return slice[n-1]
	}
	return 0
}
