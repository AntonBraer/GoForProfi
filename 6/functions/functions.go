package main

import (
	"fmt"
	"os"
	"strconv"
)

func doubleSquare(x int) (int, int) {
	return x * 2, x * x
}

func main() {
	arguments := os.Args
	if len(arguments) != 2 {
		fmt.Println("The program need 1 argument")
		return
	}

	y, err := strconv.Atoi(arguments[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	square := func(s int) int {
		return s * s
	}
	fmt.Println("The square y is", square(y))
	double := func(s int) int {
		return s + s
	}
	fmt.Println("The double y is", double(y))
	fmt.Println(doubleSquare(y))
}
