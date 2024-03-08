package piscine

func AlphaCount(s string) int {
	size := len(s)
	count := 0

	for i := 0; i < size; i++ {
		if (s[i] >= 'a' && s[i] <= 'z') || (s[i] >= 'A' && s[i] <= 'Z') {
			count++
		}
	}
	return count
}
