package main

import (
	"fmt"
	"os"
)

// Define the size of the memory and the tape array.
const memorySize = 2048

func main() {
	if len(os.Args) < 2 {
		fmt.Println("")
		return
	}

	source := os.Args[1]
	tape := make([]byte, memorySize)
	pointer := 0

	// Create a map to store the matching brackets for quick jumps
	bracketMap := make(map[int]int)

	// fmt.Println(tape)
	stack := []int{}

	// First pass: create the bracket map
	for i := 0; i < len(source); i++ {
		switch source[i] {
		case '[':
			stack = append(stack, i)
		case ']':
			if len(stack) == 0 {
				fmt.Println("Mismatched brackets")
				return
			}
			start := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			bracketMap[start] = i
			bracketMap[i] = start
		}
	}

	if len(stack) != 0 {
		fmt.Println("Mismatched brackets")
		return
	}

	// Second pass: interpret the code
	for i := 0; i < len(source); i++ {
		switch source[i] {
		case '>':
			pointer++
			if pointer >= memorySize {
				fmt.Println("Pointer out of bounds")
				return
			}
		case '<':
			pointer--
			if pointer < 0 {
				fmt.Println("Pointer out of bounds")
				return
			}
		case '+':
			tape[pointer]++
		case '-':
			tape[pointer]--
		case '.':
			fmt.Printf("%c", tape[pointer])
		case '[':
			if tape[pointer] == 0 {
				i = bracketMap[i]
			}
		case ']':
			if tape[pointer] != 0 {
				i = bracketMap[i]
			}
		}
	}
}
