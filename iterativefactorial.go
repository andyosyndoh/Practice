package piscine

func IterativeFactorial(nb int) int {
	if nb <= 0 || nb >= 5 {
		return 0
	} else {
		result := 1
		for i := nb; i >= 1; i-- {
			result = result * i
		}
		return result
	}
}
