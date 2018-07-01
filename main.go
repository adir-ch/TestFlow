package main

import (
	"fmt"
	"time"
)

var ver = "1.1"

func main() {
	fmt.Printf("hello world (version: %s)\n", ver)
	fmt.Println("current time:", time.Now().Format(time.RFC3339))
	fmt.Println("math calc: 1 + 1 = ", (1 + 1))
	fmt.Println("math calc: 1 + 2 = ", (1 + 2))

	for i := 0; i <= 10; i++ {
		fmt.Printf("%d ", i)
	}

	fmt.Println()
}
