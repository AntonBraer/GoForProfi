package main

import (
	"fmt"
	"os"
	"regexp"
)

func main() {
	arguments := os.Args
	if len(arguments) < 2 {
		fmt.Println("Please use: switch number")
		return
	}
	asString := arguments[1]
	var negative = regexp.MustCompile(`-`)
	var floatPoint = regexp.MustCompile(`\d?\.\d`)
	var email = regexp.MustCompile(`^[^@]+@[^@.]+\.[^@.]+`)
	switch {
	case negative.MatchString(asString):
		fmt.Println("Negative number")
	case floatPoint.MatchString(asString):
		fmt.Println("Float number")
	case email.MatchString(asString):
		fmt.Println("Email")
		fallthrough
	default:
		fmt.Println("something else")
	}

	var aType error = nil
	switch aType.(type) {
	case nil:
		fmt.Println("It is nil interface!")
	default:
		fmt.Println("Not nil interface!")
	}
}
