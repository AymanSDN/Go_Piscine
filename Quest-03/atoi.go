package piscine

func Atoi(s string) int {
	i := 0
	size := len(s)
	if size == 0 {
		return 0
	}
	isnegative := 1
	result := 0
	if s[i] == '-' || s[i] == '+' {
		if s[i] == '-' {
			isnegative = -1
		}
		i++
	}
	for i < size {
		if s[i] >= '0' && s[i] <= '9' {
			tmp := int(s[i] - 48)
			result = result*10 + tmp
		} else {
			return 0
		}
		i++
	}
	return result * isnegative
}
