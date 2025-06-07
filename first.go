package main

import (
	"fmt"
)

func test(num int) int {
	return num * num
}

func main() {
	fmt.Print("First home task\n")
	fmt.Println("square 5: ", test(5))
}
