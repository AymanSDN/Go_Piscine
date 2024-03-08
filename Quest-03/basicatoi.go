package piscine

func BasicAtoi(s string) int {
	var result int64 = 0

	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			tmp := int64(s[i] - 48)
			result = result*10 + tmp
		}
	}
	return int(result)
}
