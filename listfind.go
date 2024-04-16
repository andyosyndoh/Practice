package piscine

func CompStr(a, b interface{}) bool {
	return a == b
}

func ListFind(l *List, ref interface{}, comp func(a, b interface{}) bool) *interface{} {
	b := l.Head

	for b != nil {
		if comp(ref, b.Data) {
			return &b.Data
		}
		b = b.Next
	}
	return nil
}
