package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func multiplyMatrices(m1, m2 [][]int) ([][]int, error) {
	if len(m1[0]) != len(m2[0]) {
		return nil, errors.New("Cannot multiply the given matrices!")
	}
	result := make([][]int, len(m1[0]))
	for i := 0; i < len(m1); i++ {
		result[i] = make([]int, len(m2[0]))
		for j := 0; j < len(m2[0]); i++ {
			result[i][j] += m1[i][j] * m2[i][j]
		}
	}
	return result, nil
}

func createMatrix(row, col int) [][]int {
	r := make([][]int, row)
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			r[i] = append(r[i], random(-5, i*j))
		}
	}
	return r
}

func main() {
	rand.Seed(time.Now().Unix())
	arguments := os.Args
	if len(arguments) != 5 {
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
	fmt.Printf("m1 is %dx%d matrix \n", row, col)
	m1 := createMatrix(row, col)

	row, err = strconv.Atoi(arguments[3])
	if err != nil {
		fmt.Println(err)
		return
	}
	col, err = strconv.Atoi(arguments[4])
	if err != nil {
		fmt.Println(err)
		return
	}
	if col <= 0 || row <= 0 {
		fmt.Println("Need positive numbers")
		return
	}
	fmt.Printf("m2 is %dx%d matrix \n", row, col)
	m2 := createMatrix(row, col)
	fmt.Println("m1:", m1)
	fmt.Println("m2:", m2)
	r1, err := multiplyMatrices(m1, m2)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(r1)
}
