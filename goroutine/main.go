package main

import (
	"fmt"
)

var ch =  make(chan string)

func f(n int) {
	str := ""
	for i := 0; i < n; i++ {
		str += fmt.Sprintf("i=%d", i)
	}
	ch <- str
}

func main() {
	go f(10)
    go f(5)

    fmt.Println(<-ch) // Output: i=0,i=1,i=2,i=3,i=4,i=5
    fmt.Println(<-ch) // Output: i=0,i=1,i=2,i=3,i=4,i=5
}
