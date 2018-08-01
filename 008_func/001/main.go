package main

import (
	"log"
	"os"
	"strings"
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

var fm = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": firstThree,
}

func init() {
	tmp = template.Must(template.New("").Funcs(fm).ParseFiles("tmp.gohtml"))
}

func firstThree(s string) string {
	s = strings.TrimSpace(s)
	if len(s) > 3 {
		s = s[:3]
	}
	return s
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

	err := tmp.ExecuteTemplate(os.Stdout, "tmp.gohtml", data)
	if err != nil {
		log.Fatalln(err)
	}

}
