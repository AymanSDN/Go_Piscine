package piscine

func Map(f func(int) bool, a []int) []bool {
	var flags []bool
	for i := 0; i < len(a); i++ {
		flags = append(flags, f(a[i]))
	}
	return flags
}
