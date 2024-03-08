package piscine

func ShoppingListSort(slice []string) []string {
	size := len(slice)
	for i := 0; i < size; i++ {
		for j := i + 1; j < size; j++ {
			size1, size2 := len(slice[i]), len(slice[j])
			if size1 > size2 {
				slice[i], slice[j] = slice[j], slice[i]
			}
		}
	}
	return slice
}
