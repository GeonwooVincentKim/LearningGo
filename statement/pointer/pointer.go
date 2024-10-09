package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

//値レシーバ
func (p Person) BirthdayValue() {
	p.Age++
	fmt.Println("BirthdayValue p - ", p.Age)
}

// ポインタレシーバ
func (p *Person) BirthdayPointer() {
	p.Age++
	fmt.Println("BirthdayPointer p - ", p.Age)
}

// 値レシーバなのになぜReturnできる？
func (p Person) nextValue() int {
	return p.Age + 1
}

func main() {
	person := Person{Name: "ABC", Age: 30}
	fmt.Printf("First %+v\n", person)

	// こういう事を作らないようにしよう！
	// ポインタレシーバか値レシーバか判別かはっきり判別ができないからこんなコードを書かないように！
	// Go言語のルール！
	nextAge := person.nextValue()
	fmt.Println("Age - ", nextAge)
	fmt.Println("Person - ", person)
}
