package main

import (
	"log"
	"os"
	"text/template"
)

var tmp *template.Template

func init() {
	tmp = template.Must(template.ParseFiles("tmp.gohtml"))
}

func main() {
	l := []string{"eni", "beni", "vedi", "zaba", "uuu", "ddd"}
	err := tmp.Execute(os.Stdout, l)
	if err != nil {
		log.Fatalln(err)
	}
}
