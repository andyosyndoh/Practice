package piscine

func Sqrt(nb int) int {
	if nb < 0 {
		return 0
	}

	first := nb / 2

	for first := 1; first < nb; first++ {
	}
	if first*first == nb {
		return first
	} else {
		return 0
    }
}

