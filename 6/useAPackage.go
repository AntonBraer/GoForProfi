package main

import (
	"GoForProfi/6/aPackage"
	"fmt"
)

func main() {
	fmt.Println("This is main")
	aPackage.A()
	aPackage.B()
	fmt.Println(aPackage.MyConst)
}