package main

import (
	"ex04/fibonacci"
	"fmt"
)

func main() {
	arg := 4
	fmt.Println(fibonacci.Fibonacci(10))
	fmt.Println(fibonacci.Fibonacci(arg))
}
