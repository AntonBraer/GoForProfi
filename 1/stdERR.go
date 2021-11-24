package main

import (
	"io"
	"os"
)

func main() {
	myString := ""
	aruments := os.Args
	if len(aruments) == 1 {
		myString = "Give me more arguments!"
	} else {
		myString = aruments[1]
	}
	io.WriteString(os.Stdout, "Standart out!\n")
	io.WriteString(os.Stderr, myString)
	io.WriteString(os.Stdout, "\n")
}