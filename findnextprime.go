package piscine

func FindNextPrime(nb int) int {
	if nb <= 1 {
		return 2
	}
	for !IsPrime(nb) {
		nb++
	}
	return nb
}

func FindPrime(nb int) bool {
	if nb <= 1 {
		return false
	}
	if nb == 2 {
		return true
	}
	if nb%2 == 0 {
		return false
	}
	for i := 3; i*i <= nb; i += 2 {
		if nb%i == 0 {
			return false
		}
	}
	return true
}
