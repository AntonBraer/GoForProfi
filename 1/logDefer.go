package main

import (
	"fmt"
	"log"
	"os"
)

var LOGFILE = "/tmp/mGo.log"

func one(aLog *log.Logger){
	aLog.Println("-- FUNCTION ONE --")
	defer aLog.Println("-- FUNCTION ONE--")

	for i := 0; i < 10; i++ {
		aLog.Println(i)
	}
}

func two(aLog *log.Logger){
	aLog.Println("-- FUNCTION TWO --")
	defer aLog.Println("-- FUNCTION TWO--")

	for i := 10; i > 0; i-- {
		aLog.Println(i)
	}
}

func main() {
	f, err := os.OpenFile(LOGFILE, os.O_APPEND | os.O_WRONLY | os.O_CREATE, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	iLog := log.New(f, "logDefer", log.LstdFlags)
	iLog.Println("Hello!")
	one(iLog)
	two(iLog)
}