package piscine

func Abort(a, b, c, d, e int) int {
	k := []int{a,b,c,d,e}
	for a := 0; a < len(k); a++ {
		for b := a+1; b < len(k); b++ {
			if k[a]>k[b] {
				k[a],k[b] = k[b],k[a]
			}
		}
	}
	return k[2]
}
