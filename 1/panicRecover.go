package main

import "fmt"

func a() {
	fmt.Println("Inside a")
	defer func() {
		if c := recover(); c != nil{
			fmt.Println("Recover inside a")
		}
	}()
	fmt.Println("About call b")
	b()
	fmt.Println("b exited")
	fmt.Println("exited a")
}

func b() {
	fmt.Println("Inside b")
	panic("Panic in b")
	fmt.Println("exited b")
}

func main() {
	a()
	fmt.Println("main exited")
}