package piscine

import (
	"fmt"
)

func Chunk(slice []int, size int) {
	count := 0
	var new []int
	var new2 [][]int
	for i := 0; i < len(slice); i++ {
		if count < size {
			new = append(new, slice[i])
		}
		count ++ 
		if count == size || i == len(slice)-1 {
			new2 = append(new2, new)
			new = []int{}
			count = 0
		}
	}
	if size == 0 {
		fmt.Println()
	} else {
		fmt.Println(new2)
	}
	
}
