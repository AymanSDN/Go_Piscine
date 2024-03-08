package piscine

func LoafOfBread(str string) string {
	if len(str) == 0 {
		return "\n"
	} else if len(str) < 5 {
		return "Invalid Output\n"
	} else {
		s := ""
		ch := 0
		for _, char := range str {
			if ch%6 != 5 && char == ' ' {
				continue
			}
			if ch%6 == 5 {
				s += " "
			} else {
				s += string(char)
			}
			ch++
		}
		for i := len(s) - 1; i >= 0; i-- {
			if s[i] != ' ' {
				s = s[:i+1]
				break
			}
		}
		return s + "\n"
	}
}
