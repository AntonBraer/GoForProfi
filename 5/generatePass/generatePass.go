package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func main() {
	MIN := 0
	MAX := 94
	var lenght int64 = 8
	arguments := os.Args

	switch len(arguments) {
	case 2:
		lenght, _ = strconv.ParseInt(arguments[1], 10, 64)
	default:
		fmt.Println("Using def method!")
	}
	SEED := time.Now().Unix()
	rand.Seed(SEED)
	startChar := "!"
	var i int64 = 1
	for {
		myRand := random(MIN, MAX)
		newChar := string(startChar[0] + byte(myRand))
		fmt.Print(newChar)
		if i == lenght {
			break
		}
		i++
	}
	fmt.Println()
}
