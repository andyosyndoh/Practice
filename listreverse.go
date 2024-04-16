package piscine

func ListReverse(l *List) {
	var before *NodeL
	if l.Head == nil {
		return
	}
	for l.Head != nil {
		l.Head.Next, before, l.Head = before, l.Head, l.Head.Next
	}
	l.Head = before
}
