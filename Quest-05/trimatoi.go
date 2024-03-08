package piscine

func TrimAtoi(s string) int {
	s1 := []rune(s)
	var slice []rune
	isnegative := 1
	result := 0
	for i := 0; i < len(s1); i++ {
		if (s1[i] >= '0' && s1[i] <= '9') || s1[i] == '+' || s1[i] == '-' {
			slice = append(slice, s1[i])
		}
	}
	if len(slice) == 0 {
		return 0
	}
	i := 0
	if slice[i] == '-' || slice[i] == '+' {
		if slice[i] == '-' {
			isnegative *= -1
		}
		i++
	}
	for i < len(slice) {
		if slice[i] >= '0' && slice[i] <= '9' {
			tmp := int(slice[i] - 48)
			result = result*10 + tmp
		}
		i++
	}
	return result * isnegative
}
