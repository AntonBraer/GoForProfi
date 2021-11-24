package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"os"
)

func loadFromJson(filename string, key interface{}) error {
	in, err := os.Open(filename)
	if err != nil {
		return err
	}
	decodeJson := json.NewDecoder(in)
	err = decodeJson.Decode(key)
	if err != nil {
		return err
	}
	in.Close()
	return nil
}

type Record struct {
	Name    string
	Surname string
	Tel     []Telephone
}

type Telephone struct {
	Mobile bool
	Number string
}

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide file")
		return
	}
	filename := arguments[1]

	var myRecord Record
	err := loadFromJson(filename, &myRecord)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Json:", myRecord)
	myRecord.Name = "Maksim"
	xmlData, _ := xml.MarshalIndent(myRecord, "", "    ")
	xmlData = []byte(xml.Header + string(xmlData))
	fmt.Println("Xml:", string(xmlData))

	data := &Record{}
	err = xml.Unmarshal(xmlData, data)
	if err != nil {
		fmt.Println(err)
		return
	}
	result, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
		return
	}
	_ = json.Unmarshal([]byte(result), &myRecord)
	fmt.Println("json:", myRecord)
}
