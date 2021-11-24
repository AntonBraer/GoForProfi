package aPackage

import "fmt"

func A() {
	fmt.Println("This is A function")
}

func B() {
	fmt.Println("privateConstant:", privateConstant)
}

const MyConst = 5
const privateConstant = 4
