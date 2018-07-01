package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("hello world")
	fmt.Println("current time:", time.Now().Format(time.RFC3339))
}
