package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type Work struct {
	Name     string
	Category string
}

func init() {
	tpl = template.Must(template.ParseFiles("templt.gohtml"))
}

func main() {
	goDev := Work{
		Name:     "Golang Developer",
		Category: "Coding",
	}

	cyber := Work{
		Name:     "Pen Tester",
		Category: "Cyber Security",
	}

	prefWorks := []Work{goDev, cyber}

	err := tpl.Execute(os.Stdout, prefWorks)
	if err != nil {
		log.Fatalln(err)
	}
}
