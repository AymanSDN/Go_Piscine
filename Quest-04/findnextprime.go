package piscine

func Prime(n int) bool {
	if n <= 1 || n%2 == 0 || n%3 == 0 {
		return false
	}
	if n == 2 || n == 3 {
		return true
	}

	for i := 5; i*i < n; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}
	return true
}

func FindNextPrime(nb int) int {
	if Prime(nb) {
		return nb
	}
	i := 1
	for {
		if Prime(nb + i) {
			return nb + i
		}
		i++
	}
}
