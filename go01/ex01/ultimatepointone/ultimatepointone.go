package ultimatepointone

import "fmt"

func UltimatePointOne(n ***int) {
	fmt.Println("get nb value (0) -> ", n)
	fmt.Println("get nb address (0) -> \n", &n)

	fmt.Println("get nb value (1st) -> ", *n)
	fmt.Println("get nb address (1st) -> \n", &n)

	fmt.Println("get nb value (2nd) -> ", **n)
	fmt.Println("get nb address (2nd) -> \n", &n)

	***n = 1
	fmt.Println("get nb address -> ", &n)
	fmt.Println("get nb value (3rd) -> \n", ***n)
}
