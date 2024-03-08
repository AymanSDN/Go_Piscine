package piscine

func IterativePower(nb int, power int) int {
	if power < 0 {
		return 0
	}
	r := 1
	for power > 0 {
		r *= nb
		power--
	}
	return r
}
