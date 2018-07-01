package main

import (
	"fmt"
	"time"
)

var ver = "1.0"

func main() {
	fmt.Printf("hello world (version: %s)\n", ver)
	fmt.Println("current time:", time.Now().Format(time.RFC3339))
}
