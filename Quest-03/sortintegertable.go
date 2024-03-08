package piscine

func SortIntegerTable(table []int) {
	size := len(table)
	for i := 0; i < size; i++ {
		for j := 0; j < size-1-i; j++ {
			if table[j] > table[j+1] {
				tmp := table[j]
				table[j] = table[j+1]
				table[j+1] = tmp
			}
		}
	}
}