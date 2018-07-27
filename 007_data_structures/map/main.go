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
	// used for range of strings
	// idiots := []string{"Janyk", "Vova", "Pororoh", "Vata"}

	//used for map or strings
	idiots := map[string]string{
		"region": "Janyk",
		"vata":   "Vova",
		"ira":    "bera",
		"box":    "ali",
	}

	err := tmp.Execute(os.Stdout, idiots)
	if err != nil {
		log.Fatalln(err)
	}
}
