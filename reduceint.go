package piscine

import "fmt"

func ReduceInt(a []int, f func(int, int) int) {
	int1 := a[0]
	int2 := a[1]
	ans := f(int1, int2)

	fmt.Println(ans)
}
