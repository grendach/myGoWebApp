package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

var tmp *template.Template

type car struct {
	Model string
	Type  string
	Color string
}

type city struct {
	Name       string
	Location   string
	Population int
}

func init() {
	tmp = template.Must(template.ParseFiles("tmp.gohtml"))
}
func main() {
	c1 := car{
		Model: "VAZ",
		Type:  "sedan",
		Color: "black",
	}

	c2 := car{
		Model: "Mersedes",
		Type:  "cross-over",
		Color: "white",
	}

	c3 := car{
		Model: "VW",
		Type:  "sedan",
		Color: "yellow",
	}

	ct1 := city{
		Name:       "Myrgorod",
		Location:   "Ukraine",
		Population: 35000,
	}

	ct2 := city{
		Name:       "Wroclaw",
		Location:   "Poland",
		Population: 600000,
	}

	cars := []car{c1, c2, c3}
	cities := []city{ct1, ct2}

	data := struct {
		Place     []city
		Transport []car
	}{
		cities,
		cars,
	}

	err := tmp.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(cities)

}
