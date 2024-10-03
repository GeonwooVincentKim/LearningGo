package pointone

import "fmt"

func PointOne(nb *int) {
	var p *int
	fmt.Println(&p)
	fmt.Println(p)

	p = &nb
	fmt.Println(&nb)
	fmt.Println(&p)
	fmt.Println(p)
}
