package piscine

func Concat(str1 string, str2 string) string {
	var T_slice []rune
	slice1 := []rune(str1)
	slice2 := []rune(str2)
	for i := 0; i < len(slice1); i++ {
		T_slice = append(T_slice, slice1[i])
	}
	for i := 0; i < len(slice2); i++ {
		T_slice = append(T_slice, slice2[i])
	}
	return string(T_slice)
}
