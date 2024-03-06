package piscine

func FindNextPrime(nb int) int {
	if nb <= 1 {
		return 2
	} else if IsPrime(nb) {
		return nb
	} else {
		for i := nb; i <= nb+1000; i++ {
			if IsPrime(i) {
				return i
				break
			}
		}
	}
	return nb
}
