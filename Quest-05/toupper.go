package piscine

func ToUpper(s string) string {
	slice := []rune(s)
	for i := 0; i < len(slice); i++ {
		if slice[i] >= 'a' && slice[i] <= 'z' {
			slice[i] -= 32
		}
	}
	return string(slice)
}
