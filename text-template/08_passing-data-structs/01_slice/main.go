package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tmplt.gohtml"))
}

func main() {
	bizDays := []string{"Mon", "Tue", "Wed", "Thu", "Fri"}

	err := tpl.ExecuteTemplate(os.Stdout, "tmplt.gohtml", bizDays)
	if err != nil {
		log.Fatalln(err)
	}
}
