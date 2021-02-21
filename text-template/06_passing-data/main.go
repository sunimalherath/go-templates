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

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "templt.gohtml", "Friend!")
	if err != nil {
		log.Fatalln(err)
	}
}
