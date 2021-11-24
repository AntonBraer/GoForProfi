package main

import (
	"fmt"
	"log"
	"os"
)

var LOGFILE = "/tmp/mGo.log"

func main() {
	file, err := os.OpenFile(LOGFILE, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	iLog := log.New(file, "customLogLineNumber", log.LstdFlags)

	iLog.SetFlags(log.LstdFlags | log.Lshortfile)
	iLog.Println("Hello")
	iLog.Println("again")
}
