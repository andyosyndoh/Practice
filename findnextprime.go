package piscine

func FindNextPrime(nb int) int {
	if nb <= 1 {
		return nb
	} else if IsPrime(nb) {
		return nb
	} else {
    nb++
	}
	return nb
}
