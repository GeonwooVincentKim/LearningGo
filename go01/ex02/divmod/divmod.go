package divmod

import "fmt"

func DivMod(a int, b int, div *int, mod *int) {
	fmt.Println("get a value -> ", a)
	fmt.Println("get a address -> \n", &a)

	fmt.Println("get b value -> ", b)
	fmt.Println("get b address -> \n", &b)

	*div = a / b
	fmt.Println("get div value -> ", div)
	fmt.Println("get div newvalue -> ", *div)
	fmt.Println("get div address -> \n", &div)

	*mod = a % b
	fmt.Println("get mod value -> ", mod)
	fmt.Println("get mod newvalue -> ", *mod)
	fmt.Println("get mod address -> \n", &mod)
}
