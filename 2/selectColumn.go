package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	arguments := os.Args
	if len(arguments) < 2 {
		fmt.Println("usage: selectColumn column <file1> <file2>")
		return
	}
	temp, err := strconv.Atoi(arguments[1])
	if err != nil {
		fmt.Println("column is not integer")
		return
	}
	column := temp
	if column < 0 {
		fmt.Println("Column number is invalid")
		return
	}
	for _, filename := range arguments[2:] {
		fmt.Println(filename)
		f, err := os.Open(filename)
		if err != nil {
			fmt.Printf("error opening file %s\n", filename)
			continue
		}
		defer f.Close()
		r := bufio.NewReader(f)
		for {
			line, err := r.ReadString('\n')
			if err == io.EOF {
				break
			} else if err != nil {
				fmt.Printf("Error reading file %s\n", filename)
			}
			data := strings.Fields(line)
			if len(data) >= column {
				fmt.Println(data[column-1])
			}
		}
	}
}
