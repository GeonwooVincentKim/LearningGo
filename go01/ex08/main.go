package main

import (
	"ex08/basicatoi"
	"fmt"
)

func main() {
	fmt.Println(basicatoi.BasicAtoi("12345"))
	fmt.Println(basicatoi.BasicAtoi("0000000012345"))
	fmt.Println(basicatoi.BasicAtoi("000000"))
}
