package piscine

func Join(strs []string, sep string) string {
	var slice []rune
	sp := []rune(sep)
	for i := 0; i < len(strs); i++ {
		s_slice := []rune(strs[i])
		for j := 0; j < len(s_slice); j++ {
			slice = append(slice, s_slice[j])
		}
		if i < len(strs)-1 {
			for j := 0; j < len(sp); j++ {
				slice = append(slice, sp[j])
			}
		}
	}
	return string(slice)
}
