package piscine

func Rot14(s string) string {
	slice := []rune(s)
	for i := 0; i < len(slice); i++ {
		if slice[i] >= 'a' && slice[i] <= 'z' {
			slice[i] = ((slice[i] - 97 + 14) % 26) + 97
		}
		if slice[i] >= 'A' && slice[i] <= 'Z' {
			slice[i] = ((slice[i] - 65 + 14) % 26) + 65
		}
	}
	return string(slice)
}
