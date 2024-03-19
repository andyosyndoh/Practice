package piscine

func ListClear(l *List) {
	currentNode := l.Head

	for currentNode != nil {
		nextNode := currentNode.Next

		l.Head = nextNode

		currentNode = nextNode
	}

	l.Head = nil
	l.Tail = nil
}
