package piscine

func Split(str string) []string {
	var splitted []string
	var w string
	for i := 0; i < len(str); i++ {
		if str[i] != ' ' {
			w += string(str[i])
		} else {
			splitted = append(splitted, w)
			w = ""
		}
	}
	if w != " " {
		splitted = append(splitted, w)
	}
	return splitted
}

func ShoppingSummaryCounter(str string) map[string]int {
	word := Split(str)
	sum := make(map[string]int)
	for i := 0; i < len(word); i++ {
		sum[word[i]]++
	}
	return sum
}
