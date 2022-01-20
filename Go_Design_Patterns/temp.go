package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"text/template"
)

type Book struct {
	Name   string
	Author string
}

type IBook interface {
	addBook(b Book)
	addAuthor(b Book)
}

func main() {
	b1 := Book{"BhagavadGita", "Prabhupada"}
	temp_b1 := template.New("Template_1")
	//Author := "agawane"
	temp_b1, _ = temp_b1.Parse(" {{.Name}}" + "_" + " .mpd" + "{{.Author}}")
	// standard output to print merged data
	var tpl bytes.Buffer
	//_ = temp_b1.Execute(os.Stdout, b1)
	if err := temp_b1.Execute(&tpl, IBook.addBook(b1)); err != nil {
		log.Fatalf("failed to parse")
	}
	fmt.Println("Printing with String method")
	fmt.Println(tpl.String())

}
