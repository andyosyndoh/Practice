package piscine

func FindNextPrime(nb int) int {
	if nb <= 1 {
		return nb
	} else if IsPrime(nb) {
		return nb
	} else {
		for i := nb; i <= nb+3; i++{
			if IsPrime(i) {
				return i
			}
		}
	}
	return nb
}
