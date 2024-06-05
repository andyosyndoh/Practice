package main

import (
	"fmt"
	"os"
	"piscine"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println(0)
		return
	}
	num := piscine.BasicAtoi2(os.Args[1])
	if num < 1 {
		fmt.Println(0)
		return
	}

	ans := 0
	for i := 1; i <= num; i++ {
		if piscine.IsPrime(i) {
			ans += i
		}
	}
	fmt.Println(ans)
}
