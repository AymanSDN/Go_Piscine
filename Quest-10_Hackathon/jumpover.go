package piscine

func JumpOver(str string) string {
	var slice []rune
	size := len(str)
	if size == 0 || size < 2 {
		return "\n"
	}
	for i := 2; i < size; i += 3 {
		slice = append(slice, rune(str[i]))
	}
	return string(slice) + "\n"
}
