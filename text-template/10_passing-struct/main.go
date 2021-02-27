package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("templt.gohtml"))
}

type Subject struct {
	Name   string
	Rating int
}

func main() {
	golang := Subject{
		Name:   "Golang",
		Rating: 1,
	}

	err := tpl.Execute(os.Stdout, golang)
	if err != nil {
		log.Fatalln(err)
	}
}
