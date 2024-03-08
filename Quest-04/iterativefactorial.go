package piscine

func IterativeFactorial(nb int) int {
	result := 1
	if nb > 20 || nb < 0 {
		return 0
	}
	for i := nb; i > 1; i-- {
		result *= i
	}
	return result
}
