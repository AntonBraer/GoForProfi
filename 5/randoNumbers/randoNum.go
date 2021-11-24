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
	MAX := 100
	TOTAL := 100
	seed := time.Now().Unix()
	arguments := os.Args

	switch len(arguments) {
	case 2:
		fmt.Println("Usage ./rndomNum MIN MAX TOTAL SEED")
		MIN, _ = strconv.Atoi(arguments[1])
		MAX = MIN + 100
	case 3:
		fmt.Println("Usage ./rndomNum MIN MAX TOTAL SEED")
		MIN, _ = strconv.Atoi(arguments[1])
		MAX, _ = strconv.Atoi(arguments[2])
	case 4:
		fmt.Println("Usage ./rndomNum MIN MAX TOTAL SEED")
		MIN, _ = strconv.Atoi(arguments[1])
		MAX, _ = strconv.Atoi(arguments[2])
		TOTAL, _ = strconv.Atoi(arguments[3])
	case 5:
		fmt.Println("Usage ./rndomNum MIN MAX TOTAL SEED")
		MIN, _ = strconv.Atoi(arguments[1])
		MAX, _ = strconv.Atoi(arguments[2])
		TOTAL, _ = strconv.Atoi(arguments[3])
		seed, _ = strconv.ParseInt(arguments[4], 10, 64)
	default:
		fmt.Println("Def")
	}
	rand.Seed(seed)
	for i := 0; i < TOTAL; i++ {
		myrand := random(MIN, MAX)
		fmt.Print(myrand, " ")
	}
	fmt.Println()
}
