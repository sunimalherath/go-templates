package main

import (
	"strings"
	"text/template"
)

var tpl *template.Template

var fm = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": firstThree,
}

func init() {

}

func firstThree(s string) string {
	s = strings.TrimSpace(s)
	s = s[0:3]
	return s
}

func main() {

}
