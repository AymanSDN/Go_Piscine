package piscine

func Unmatch(a []int) int {
	size := len(a)
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if !(i == j || a[i] != a[j]) {
				a[i], a[j] = -1, -1
			}
		}
	}
	for i := 0; i < size; i++ {
		if a[i] != -1 {
			return a[i]
		}
	}
	return -1
}
