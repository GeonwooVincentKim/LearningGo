package swap

import "fmt"

func Swap(a *int, b *int) {
	tmp := *a

	fmt.Println("Get A Address -> ", &a)
	fmt.Println("Get A Value -> \n", a)

	fmt.Println("Get Tmp Address -> ", &tmp)
	fmt.Println("Get Tmp Value -> \n", tmp)

	*a = *b
	fmt.Println("Get A Address -> ", &a)
	fmt.Println("Get A Value -> \n", a)

	fmt.Println("Get B Address -> ", &b)
	fmt.Println("Get B Value -> \n", b)

	*b = tmp
	fmt.Println("Get B Address -> ", &b)
	fmt.Println("Get B Value -> \n", b)

	fmt.Println("Get Tmp Address -> ", &tmp)
	fmt.Println("Get Tmp Value -> \n", tmp)
}
