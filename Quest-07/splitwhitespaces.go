package piscine

func SplitWhiteSpaces(s string) []string {
	slice := []rune(s)
	var s1 []string
	var tmp string
	i := 0
	for i < len(slice) {
		if !(slice[i] == ' ') {
			tmp += string(slice[i])
			i++
		} else {
			if tmp != " " {
				s1 = append(s1, tmp)
			}
			tmp = ""
			for i < len(slice) && (slice[i] == ' ' || slice[i] == '\n' || slice[i] == '\t') {
				i++
			}
		}
	}
	s1 = append(s1, tmp)
	return s1
}
