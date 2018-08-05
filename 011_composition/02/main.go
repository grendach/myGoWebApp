package main

import (
	"log"
	"os"
	"text/template"
)

var tmp *template.Template

type person struct {
	Name string
	Age  int
}

func init() {
	tmp = template.Must(template.ParseFiles("tmp.gohtml"))
}

func main() {
	p1 := person{
		Name: "Dmytro Grendach",
		Age:  33,
	}

	err := tmp.Execute(os.Stdout, p1)
	if err != nil {
		log.Fatalln(err)
	}

}
