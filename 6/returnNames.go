package main

import (
	"fmt"
	"os"
	"strconv"
)

func namedMinMax(x, y int) (min, max int) {
	if x > y {
		min = y
		max = x
	} else {
		min = x
		max = y
	}
	return
}

func main() {
	arguments := os.Args
	if len(arguments) < 3 {
		fmt.Println("The program needs at least 2 arguments!")
		return
	}
	a1, _ := strconv.Atoi(arguments[1])
	a2, _ := strconv.Atoi(arguments[2])
	fmt.Println(namedMinMax(a1, a2))

}
