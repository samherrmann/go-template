package main

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"os"
)

// JSON ...
type JSON map[string]interface{}

func main() {

	os.MkdirAll("dist", os.ModePerm)
	compileTemplate(readJSONFile("data.json"))
}

func readJSONFile(fileName string) JSON {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	buffer, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	data := make(JSON)
	if err := json.Unmarshal(buffer, &data); err != nil {
		panic(err)
	}
	return data
}

func compileTemplate(data interface{}) {
	fileName := "index.html"
	tpl, err := template.ParseFiles(fileName)
	if err != nil {
		panic(err)
	}
	file, err := os.Create("dist/" + fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	if err := tpl.Execute(file, data); err != nil {
		panic(err)
	}
}
