package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

type Entry struct {
	Number int
	Double int
	Square int
}

var DATA []Entry
var tFile string

func myHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Host: %s Path: %s\n", r.Host, r.URL.Path)
	myT := template.Must(template.ParseGlob(tFile))
	myT.ExecuteTemplate(w, tFile, DATA)
}

func main() {
	arguments := os.Args
	if len(arguments) != 2 {
		fmt.Println("Need template file!")
		return
	}
	tFile = arguments[1]
	for i := 0; i < 10; i++ {
		temp := Entry{Number: i, Double: i + i, Square: i * i}
		DATA = append(DATA, temp)
	}

	http.HandleFunc("/", myHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}
