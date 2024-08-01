package main

import (
	"fmt"
	"piscine"
)

func main() {
	input1 := []uint{2, 3, 1, 1, 4}
	fmt.Println(piscine.CanJump(input1))

	input2 := []uint{3, 2, 1, 0, 4}
	fmt.Println(piscine.CanJump(input2))

	input3 := []uint{0}
	fmt.Println(piscine.CanJump(input3))

	input4 := []uint{0, 0, 0, 0, 0, 0}
	fmt.Println(piscine.CanJump(input4))

	input5 := []uint{1,2,3,4,5,0}
	fmt.Println(piscine.CanJump(input5))

	input6 := []uint{3,2,0,1,4}
	fmt.Println(piscine.CanJump(input6))
}
