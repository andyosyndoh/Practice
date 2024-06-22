package main

import (
	"fmt"
	"piscine"
)

func main() {
	s := "HelloHAhowHAareHAyou?"
	result := piscine.Split(s, "HA")
	fmt.Printf("%#v\n", result)
}
