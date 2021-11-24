package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"os"
	"strconv"
)

func generateBytes(n int64) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}
	return b, err
}

func generatePass(s int64) (string, error) {
	b, err := generateBytes(s)
	return base64.URLEncoding.EncodeToString(b), err
}

func main() {
	var length int64 = 8
	arguments := os.Args
	switch len(arguments) {
	case 2:
		length, _ = strconv.ParseInt(arguments[1], 10, 64)
		if length == 0 {
			length = 8
		}
	}
	myPass, err := generatePass(length)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(myPass[0:length])
}
