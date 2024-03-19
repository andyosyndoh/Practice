package piscine

func ListAt(l *NodeL, pos int) *NodeL {
    if l == nil {
        return nil
    }

    current := l
    currentPos := 0
    for current != nil {
        if currentPos == pos {
            return current
        }
        currentPos++
        current = current.Next
    }
    return nil
}
