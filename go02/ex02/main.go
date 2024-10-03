package main

import (
	"ex02/iterativepower"
	"fmt"
)

func test() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}

func main() {
	fmt.Println(test())

	fmt.Println(iterativepower.IterativePower(5, 3))
	fmt.Println(iterativepower.IterativePower(4, 3))
}
