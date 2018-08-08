package main

import (
	"log"
	"os"
	"text/template"
)

type person struct {
	Name string
	Age  int
}

func (p person) SomeProcess() int {
	return 7
}

func (p person) AgeDbl() int {
	return p.Age * 2
}

func (p person) AgeThree(x int) int {
	return x * 2
}

var tmp *template.Template

func init() {
	tmp = template.Must(template.ParseFiles("tmp.gohtml"))
}

func main() {
	p := person{
		Name: "Dmytro",
		Age:  33,
	}

	err := tmp.Execute(os.Stdout, p)
	if err != nil {
		log.Fatalln(err)
	}
}
