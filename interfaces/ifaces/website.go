package ifaces

import "fmt"

type Web interface {
	GetId() int
	GetName() string
}

type Amazon struct {
	ProductId       int
	ProductContents string
	Location        string
}

type Google struct {
	UserId   int
	UserName string
}

type Naver struct {
	AdvancedMap string
}

func (a Amazon) GetId() int {
	fmt.Println("Your Amazon Account ID is ", a.ProductId)
	return a.ProductId
}

func (a Amazon) GetName() string {
	fmt.Println("Your Amazon Account Name is ", a.ProductContents)
	return a.ProductContents
}

func (g Google) GetId() int {
	fmt.Println("Your Google Account ID is ", g.UserId)
	return g.UserId
}

func (g Google) GetName() string {
	fmt.Println("Your Google Account Name is ", g.UserName)
	return g.UserName
}

func (n Naver) GetId() int {
	fmt.Println("Your Naver Account ID is ", 1)
	return 1
}

func (n Naver) GetName() string {
	fmt.Println("Your Naver Account Name is ", n.AdvancedMap)
	return n.AdvancedMap
}

func GetAllInfo(w Web) {
	w.GetId()
	w.GetName()
}
