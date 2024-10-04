package ultidivmod

import "fmt"

func UltiDivMod(a *int, b *int) {
	fmt.Println("Address uA -> ", &a)
	fmt.Println("Value uA -> \n", a)

	fmt.Println("Address uB -> ", &b)
	fmt.Println("Value uB -> \n", b)

	div := *a / *b
	fmt.Println("Address uDiv -> ", &div)
	fmt.Println("Value uDiv -> \n", div)

	mod := *a % *b
	fmt.Println("Address uMod -> ", &mod)
	fmt.Println("Value uuModDiv -> \n", mod)

	*a = div
	fmt.Println("Address uA -> ", &a)
	fmt.Println("Value uA -> \n", a)

	*b = mod
	fmt.Println("Address uB -> ", &b)
	fmt.Println("Value uB -> \n", b)
}
