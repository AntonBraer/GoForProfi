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

func negativeMatrices(s [][]int) [][]int {
	for i, x := range s {
		for j, _ := range x {
			s[i][j] = -s[i][j]
		}
	}
	return s
}

func addMartices(m1, m2 [][]int) [][]int {
	result := make([][]int, len(m1))
	for i, x := range m1 {
		for j, _ := range x {
			result[i] = append(result[i], m1[i][j]+m2[i][j])
		}
	}
	return result
}

func main() {
	arguments := os.Args
	if len(arguments) != 3 {
		fmt.Println("Wrong numbers of arguments")
		return
	}
	var row, col int
	row, err := strconv.Atoi(arguments[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	col, err = strconv.Atoi(arguments[2])
	if err != nil {
		fmt.Println(err)
		return
	}
	if col <= 0 || row <= 0 {
		fmt.Println("Need positive numbers")
		return
	}
	fmt.Printf("Using %dx%d arrays \n", row, col)

	m1 := make([][]int, row)
	m2 := make([][]int, row)
	rand.Seed(time.Now().Unix())

	for i := 0; i < row; i++ {
		for j := 1; j < col; j++ {
			m1[i] = append(m1[i], random(-1, i*j+rand.Intn(10)))
			m2[i] = append(m2[i], random(-1, i*j+rand.Intn(10)))
		}
	}
	fmt.Println("m1:", m1)
	fmt.Println("m2:", m2)
	r1 := addMartices(m1, m2)
	r2 := addMartices(m1, negativeMatrices(m2))
	fmt.Println("r1:", r1)
	fmt.Println("r2:", r2)

}
