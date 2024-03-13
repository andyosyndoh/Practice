package piscine

func Unmatch(a []int) int {
	new := make(map[int]int)
	for _, char := range a {
		new[char]++
	}
	minUnmatched := -1
    for char, count := range new {
        if count == 1 {
            if minUnmatched == -1 || char < minUnmatched {
                minUnmatched = char
            }
        }
    }
    
    return minUnmatched
}

