package piscine

// This function reverse a string in place and return it

func StrRev(s string) string {
	slice := []rune(s)
	size := len(slice) - 1

	for i := 0; i < size; i++ {
		tmp := slice[i]
		slice[i] = slice[size]
		slice[size] = tmp
		size--
	}
	return string(slice)
}

// func main() {
// 	s := "Hello World!"
// 	s = StrRev(s)
// 	fmt.Println(s)
// }
