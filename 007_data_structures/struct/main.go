package main

import (
	"log"
	"os"
	"text/template"
)

var tmp *template.Template

type transport struct {
	Model string
	Type  string
}

func init() {
	tmp = template.Must(template.ParseFiles("tmp.gohtml"))
}

func main() {
	myCar := transport{
		Model: "Fabia",
		Type:  "car",
	}
	myMoto := transport{
		Model: "BMW",
		Type:  "motobyke",
	}

	myBicycle := transport{
		Model: "B-Twin",
		Type:  "bicycle",
	}

	tr := []transport{myBicycle, myCar, myMoto}

	err := tmp.Execute(os.Stdout, tr)
	if err != nil {
		log.Fatalln(err)
	}
}
