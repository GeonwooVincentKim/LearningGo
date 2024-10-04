package pointone

import "fmt"

func PointOne(nb *int) {
	fmt.Println("get nb address -> ", &nb)
	fmt.Println("get nb value -> ", nb)
	fmt.Println()
	*nb = 1
	fmt.Println("get nb asterisk -> ", *nb)
	fmt.Println("get nb address -> ", &nb)
	fmt.Println("get nb value -> ", nb)
}
