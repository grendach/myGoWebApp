package main

import (
	"log"
	"os"
	"text/template"
)

type hotel struct {
	Name, Address, City, Region, Zip string
}

var tmp *template.Template

func init() {
	tmp = template.Must(template.ParseFiles("tmp.gohtml"))
}

func main() {
	h := hotel{
		Name:    "Lybid",
		Address: "Khreshchatyk 18",
		City:    "Kyiv",
		Region:  "Ceter",
		Zip:     "00440",
	}
	err := tmp.Execute(os.Stdout, h)
	if err != nil {
		log.Println(err)
	}
}
