package main

import (
	"fmt"
	"os"
)

func main() {
	input := os.Args[1:]
	for _, char := range input {
		if char == "01" || char == "galaxy" || char == "galaxy 01" {
			fmt.Println("Alert!!!")
			return
		}
	}
}
