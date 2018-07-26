package main

import (
	"log"
	"os"
	"text/template"
)

var tlp *template.Template

func init() {
	tlp = template.Must(template.ParseFiles("hello.gohtml"))
}
func main() {
	word := tlp.ExecuteTemplate(os.Stdout, "hello.gohtml", 33)
	if word != nil {
		log.Fatalln(word)
	}

}
