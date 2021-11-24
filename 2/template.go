package main

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
)

type dataView struct {
	h1 string
	h2 string
}

func main() {
	f, err := os.Open("1.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	b, err := io.ReadAll(f)
	if err != nil {
		fmt.Println(err)
		return
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := dataView{
			h1: "hello",
			h2: "world",
		}
		tmpl, _ := template.New("data").Parse(string(b))
		tmpl.Execute(w, data)
	})

	fmt.Println("Server is listening...")
	http.ListenAndServe(":8080", nil)
}
