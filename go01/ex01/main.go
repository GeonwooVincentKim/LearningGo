package main

import (
	"ex01/ultimatepointone"
	"fmt"
)

func main() {
	a := 0
	fmt.Println("get a value -> ", a)
	fmt.Println("get a address -> \n", &a)

	b := &a
	fmt.Println("get b value -> ", b)
	fmt.Println("get b address -> \n", &b)

	c := &b
	fmt.Println("get c value -> ", c)
	fmt.Println("get c address -> \n", &c)

	ultimatepointone.UltimatePointOne(&c)
	fmt.Println(c)
}
