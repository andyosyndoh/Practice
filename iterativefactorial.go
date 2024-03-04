package piscine

func IterativeFactorial(nb int) int {
	if nb < 0 || nb >  20 {
		return 0
	} else if nb == 0 {
		return 1
	} else {
		result := 1
		for i := nb; i >= 1; i-- {
			result = result * i
		}
		return result
	}
}
