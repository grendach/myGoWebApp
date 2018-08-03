package main

import (
	"log"
	"os"
	"text/template"
)

var tmp *template.Template

func init() {
	tmp = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	err := tmp.ExecuteTemplate(os.Stdout, "index.gohtml", "bugaga")
	if err != nil {
		log.Println(err)
	}
}
