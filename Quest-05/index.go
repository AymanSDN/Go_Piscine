package piscine

func Index(s string, toFind string) int {
	s_size := len(s)
	tF_size := len(toFind)

	for i := 0; i < s_size; i++ {
		if s[i] == toFind[0] {
			for j := 0; j < tF_size; j++ {
				if s[i+j] != toFind[j] {
					break
				}
				if j == tF_size-1 {
					return i
				}
			}
		}
	}
	return -1
}
