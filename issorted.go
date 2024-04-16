package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	length := 0
	for i := range a {
		length = i + 1
	}

	asciending := true
	descending := true

	for i := 1; i < length; i++ {
		if !(f(a[i-1], a[i]) >= 0) {
			descending = false
		}
	}

	for i := 1; i < length; i++ {
		if !(f(a[i-1], a[i]) <= 0) {
			asciending = false
		}
	}

	return asciending || descending
}
