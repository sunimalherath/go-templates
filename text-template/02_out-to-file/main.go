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

	nf, err := os.Create("index.html")
	if err != nil {
		log.Fatalln(err)
	}
	defer nf.Close()

	err = tmplt.Execute(nf, nil)
	if err != nil {
		log.Fatalln(err)
	}
}
