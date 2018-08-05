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

type JamesBond struct {
	person
	LicenseToKill bool
}

func init() {
	tmp = template.Must(template.ParseFiles("tmp.gohtml"))
}

func main() {
	p1 := JamesBond{
		person{
			Name: "Ilon Musk",
			Age:  46,
		},
		true,
	}

	err := tmp.Execute(os.Stdout, p1)
	if err != nil {
		log.Fatalln(err)
	}

}
