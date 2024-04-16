package piscine

func ListSize(l *List) int {
	count := 0
	new := l.Head
	for new != nil {
		count++
		new = new.Next
	}
	return count
}
