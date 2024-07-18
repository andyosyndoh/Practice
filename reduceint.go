package piscine

import (
	"github.com/01-edu/z01"
	"strconv"
)

func ReduceInt(a []int, f func(int, int) int) {
	ans := 0

	int1 := a[0]
	int2 := a[1]

	ans = f(int1, int2)

	ans2 := strconv.Itoa(ans)

	for _ , ch := range ans2 {
		z01.PrintRune(ch)
	}
	z01.PrintRune('\n')
}
