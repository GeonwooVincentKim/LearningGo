package main

import (
	"ex03/ultidivmod"
	"fmt"
)

func main() {
	a := 13
	b := 2

	fmt.Println("get a value -> ", a)
	fmt.Println("get a address -> \n", &a)
	fmt.Println("get b value -> ", b)
	fmt.Println("get b address -> \n", &b)

	ultidivmod.UltiDivMod(&a, &b)
	fmt.Println(a)
	fmt.Println(b)
}
