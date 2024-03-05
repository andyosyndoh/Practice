package piscine

func Sqrt(nb int) int {
	if nb == 0 {
		return nb
	}
	if nb == 1 {
		return nb
	}
	first := nb / 2
	for i := 1; i < nb; i++ {
		if first != 0 {
			first = (first + nb/first) / 2
		} else {
			break
		}
	}
	if first*first != nb {
		return 0
	}
	return first
}
