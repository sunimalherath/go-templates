package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	/*
		// template.ParseFiles return a pointer to a Template - in this case "tplt"
		// "tplt" is can be think as a collection (collects templates)
	*/
	tplt, err := template.ParseFiles("t1.dat")
	if err != nil {
		log.Fatalln(err)
	}

	err = tplt.Execute(os.Stdout, nil)

	tplt, err = tplt.ParseFiles("t2.text", "t3.text")
	if err != nil {
		log.Fatalln(err)
	}

	// ExecuteTemplate - helps to run a specific template from the collection of templates.
	err = tplt.ExecuteTemplate(os.Stdout, "t2.text", nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = tplt.ExecuteTemplate(os.Stdout, "t3.text", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
