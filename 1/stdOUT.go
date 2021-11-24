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
	io.WriteString(os.Stdout, myString)
	io.WriteString(os.Stdout, "\n")
}