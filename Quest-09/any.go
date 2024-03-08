package piscine

func Any(f func(string) bool, a []string) bool {
	flag := false
	for i := 0; i < len(a); i++ {
		if f(a[i]) {
			flag = f(a[i])
		}
	}
	return flag
}
