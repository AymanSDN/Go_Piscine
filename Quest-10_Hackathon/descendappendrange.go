package piscine

func DescendAppendRange(max, min int) []int {
	tab := []int{}
	for ; max > min; max-- {
		tab = append(tab, max)
	}
	return tab
}
