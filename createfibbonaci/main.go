package main

import "fmt"


func main() {
	nums := []int{0,1}
	for nums[len(nums)-1] < 100 {
		nums = append(nums, nums[len(nums)-1]+nums[len(nums)-2])
	}
	fmt.Println(nums)
}