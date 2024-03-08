package piscine

func BasicAtoi2(s string) int {
	var result int = 0

	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			tmp := int(s[i] - 48)
			result = result*10 + tmp
		} else {
			return 0
		}
	}
	return int(result)
}