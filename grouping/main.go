package main

import (
	"fmt"
	"os"
	"strings"
)


func main(){
	if len(os.Args) != 3 {
		return
	}
	str1 := os.Args[1]
	str2 := os.Args[2]

	if len(str1) < 3 || len(str2) < 1 {
		return
	}

	if str1[0] != '(' || str1[len(str1)-1] != ')' {
		return
	}

	regs := strings.Split(str1[1:len(str1)-1], "|")

	words := strings.Split(str2, " ")

	count := 0
	for _, word := range words {
		for _, reg := range regs{
			if strings.Contains(word, reg) {
				count++
				if word[len(word)-1] == ','{
					word = word[:len(word)-1]
				}
				fmt.Printf("%v: %v\n", count,word)
			}
		}
	}
}