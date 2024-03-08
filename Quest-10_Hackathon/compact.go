package piscine

func Compact(ptr *[]string) int {
	tab := *ptr
	var slice []string

	for i := 0; i < len(tab); i++ {
		if tab[i] != "" {
			slice = append(slice, tab[i])
		}
	}
	*ptr = slice
	return len(slice)
}
