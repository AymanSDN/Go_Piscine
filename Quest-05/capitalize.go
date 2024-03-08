package piscine

func isalphanum(s string) bool {
	for i := 0; i < len(s); i++ {
		if (s[i] < 'A' || s[i] > 'Z') && (s[i] < 'a' || s[i] > 'z') && (s[i] < '0' || s[i] > '9') {
			return false
		}
	}
	return true
}

func tolow(s string) string {
	slice := []rune(s)
	for i := 0; i < len(slice); i++ {
		if slice[i] >= 'A' && slice[i] <= 'Z' {
			slice[i] += 32
		}
	}
	return string(slice)
}

func Capitalize(s string) string {
	s = tolow(s)
	s1 := []rune(s)
	if isalphanum(string(s1[0])) {
		if s1[0] >= 'a' && s1[0] <= 'z' {
			s1[0] -= 32
		}
	}
	for i := 1; i < len(s1); i++ {
		if !isalphanum(string(s1[i-1])) {
			if s1[i] >= 'a' && s1[i] <= 'z' {

				s1[i] -= 32
			}
		}
	}
	return string(s1)
}
