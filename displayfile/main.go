package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("File name missing")
		return
	} else if len(args) > 1 {
		fmt.Println("Too many arguments")
		return
	}

	fileName := args[0]

	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fmt.Print(string(data))
}
