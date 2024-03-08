package piscine

func ToLower(s string) string {
	slice := []rune(s)
	for i := 0; i < len(slice); i++ {
		if slice[i] >= 'A' && slice[i] <= 'Z' {
			slice[i] += 32
		}
	}
	return string(slice)
}
