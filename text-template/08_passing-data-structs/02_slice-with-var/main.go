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
	bizDays := []string{"Mon", "Tue", "Wed", "Thu", "Fri"}

	err := tpl.ExecuteTemplate(os.Stdout, "templt.gohtml", bizDays)
	if err != nil {
		log.Fatalln(err)
	}
}
