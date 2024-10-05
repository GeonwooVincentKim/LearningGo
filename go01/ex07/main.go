package main

import (
	"ex07/strrev"
	"fmt"
)

func main() {
	s := "Hello World!"
	s = strrev.StrRev(s)
	fmt.Println(s)
}
