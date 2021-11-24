package main

import "fmt"

func main() {
	c1 := 12 + 1i
	c2 := complex(5, 7)
	fmt.Println(c1)
	fmt.Println(c2)
	var c3 complex64 = complex64(c1 + c2)
	fmt.Println(c3)
	cZero := c3 - c3
	fmt.Printf("%T", cZero)

}
