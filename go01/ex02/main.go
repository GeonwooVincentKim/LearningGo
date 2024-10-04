package main

import (
	"ex02/divmod"
	"fmt"
)

func main() {
	a := 13
	b := 2

	fmt.Println("get a value -> ", a)
	fmt.Println("get a address -> \n", &a)
	fmt.Println("get b value -> ", b)
	fmt.Println("get b address -> \n", &b)

	var div int
	var mod int
	divmod.DivMod(a, b, &div, &mod)
	fmt.Println(div)
	fmt.Println(mod)
}
