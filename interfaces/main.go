package main

import (
	"fmt"
	"interfaces/ifaces"
)

func main() {
	var amazonAccount1 = ifaces.Amazon{ProductId: 1, ProductContents: "User", Location: "Somewhere"}
	var googleAccount2 = ifaces.Google{UserId: 2, UserName: "User2"}
	var naverAccount3 = ifaces.Naver{AdvancedMap: "New Map"}

	var getWeb ifaces.Web
	fmt.Println("Testing Amazon")
	getWeb = amazonAccount1
	ifaces.GetAllInfo(getWeb)

	fmt.Println("Testing Google")
	getWeb = googleAccount2
	ifaces.GetAllInfo(getWeb)

	fmt.Println("Testing Naver")
	getWeb = naverAccount3
	ifaces.GetAllInfo(getWeb)
}
