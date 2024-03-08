package piscine

func StringToIntSlice(str string) []int {
	slice := []rune(str)
	var tab []int
	for i := 0; i < len(slice); i++ {
		tab = append(tab, int(slice[i]))
	}
	return tab
}
