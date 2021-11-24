package main

// #include <stdio.h>
// void callC() {
//		printf("Call!\n");
// }
import "C"
import "fmt"

func main() {
	fmt.Println("GO!")
	C.callC()
	fmt.Println("Again GO!s")
}