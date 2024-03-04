package piscine

func IterativePower(nb int, power int) int {
	if power < 0 || power > 20 {
		return 0
	} else {
		result := 1
		for i := 1; i <= power; i++ {
			result = result * nb
		}
		return result
	}
}
