package piscine

func MakeRange(min, max int) []int {
	if max < min || (min == 0 && max == 0) {
		return nil
	}
	size := max - min
	tab := make([]int, size)
	for i := 0; i < size; i++ {
		tab[i] = min
		min++
	}
	return tab
}