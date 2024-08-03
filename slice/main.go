package main

import (
	"fmt"
)

func main() {
	a := []string{"coding", "algorithm", "ascii", "package", "golang"}
	fmt.Printf("%#v\n", Slice(a, 1))
	fmt.Printf("%#v\n", Slice(a, 2, 4))
	fmt.Printf("%#v\n", Slice(a, -3))
	fmt.Printf("%#v\n", Slice(a, -2, -1))
	fmt.Printf("%#v\n", Slice(a, 2, 0))
}

func Slice(a []string, nbrs ...int) []string {
	if len(nbrs) == 1 && nbrs[0] > 0 {
		return a[nbrs[0]:]
	}
	if len(nbrs) == 1 && nbrs[0] < 0 {
		i := len(a) + nbrs[0]
		return a[i:]
	}
	if nbrs[0] > 0 && nbrs[1] > 0 {
		return a[nbrs[0]:nbrs[1]]
	}
	if nbrs[0] < 0 && nbrs[1] < 0 {
		start := len(a) + nbrs[0]
		finish := len(a) + nbrs[1]
		return a[start:finish]
	}

	return nil
}
