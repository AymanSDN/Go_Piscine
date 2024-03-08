package piscine

func BasicJoin(elems []string) string {
	var slice []rune

	for i := 0; i < len(elems); i++ {
		s_slice := []rune(elems[i])
		for j := 0; j < len(s_slice); j++ {
			slice = append(slice, s_slice[j])
		}
	}
	return string(slice)
}
