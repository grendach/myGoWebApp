package main

import (
	"log"
	"os"
	"text/template"
)

var tmp *template.Template

type car struct {
	Model     string
	Producent string
}

func init() {
	tmp = template.Must(template.ParseFiles("tmp.gohtml"))
}

func main() {
	myCar := car{
		Model:     "Fabia",
		Producent: "Scoda",
	}

	err := tmp.Execute(os.Stdout, myCar)
	if err != nil {
		log.Fatalln(err)
	}
}
