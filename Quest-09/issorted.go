package piscine

func f(a, b int) int {
	return a - b
}

func IsSorted(f func(a, b int) int, a []int) bool {
	desc := true
	asc := true
	for i := 1; i < len(a); i++ {
		if f(a[i], a[i-1]) < 0 {
			desc = false
		}
		if f(a[i], a[i-1]) > 0 {
			asc = false
		}
	}
	return asc || desc
}
