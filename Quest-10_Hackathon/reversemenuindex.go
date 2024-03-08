package piscine

func ReverseMenuIndex(menu []string) []string {
	size := len(menu)
	rev := make([]string, size)
	for i := 0; i < size; i++ {
		rev[size-1-i] = menu[i]
	}
	return rev
}
