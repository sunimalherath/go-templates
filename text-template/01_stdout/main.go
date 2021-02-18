package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	tmplt, err := template.ParseFiles("tmplt.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	err = tmplt.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}
}
