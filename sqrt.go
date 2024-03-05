package piscine

func Sqrt(nb int) int {
	if nb < 0 {
		return 0
	} else if nb == 1 {
		return 1
	}

	first := nb / 2

	for first := 2; first < nb; first++ {
	}
	if first*first == nb {
		return first
	} else {
		return 0
	}
}
