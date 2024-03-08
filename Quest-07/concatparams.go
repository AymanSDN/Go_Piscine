package piscine

func ConcatParams(args []string) string {
	size := len(args)
	var slice []rune
	for i := 0; i < size; i++ {
		for j := 0; j < len(args[i]); j++ {
			slice = append(slice, rune(args[i][j]))
		}
		if i < size-1 {
			slice = append(slice, '\n')
		}
	}
	return string(slice)
}
