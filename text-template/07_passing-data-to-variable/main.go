package main

import (
	"log"
	"os"
	"text/template"
)

var tplt *template.Template

func init() {
	tplt = template.Must(template.ParseFiles("templt.gohtml"))
}

func main() {
	err := tplt.ExecuteTemplate(os.Stdout, "templt.gohtml", `Hello friend, how are things in Go world?`)
	if err != nil {
		log.Fatalln(err)
	}
}
