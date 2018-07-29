package main

import (
	"log"
	"os"
	"text/template"
)

var tmp *template.Template

type tree struct {
	Color string
	Type  string
}

type car struct {
	Model string
	Type  string
	Color string
}

type items struct {
	Plant     []tree
	Transport []car
}

func init() {
	tmp = template.Must(template.ParseFiles("tmp.gohtml"))
}
func main() {
	t1 := tree{
		Type:  "apple",
		Color: "green",
	}

	t2 := tree{
		Type:  "cherry",
		Color: "red",
	}

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

	trees := []tree{t1, t2}
	cars := []car{c1, c2}

	// data1 := items{
	// 	Plant:     trees,
	// 	Transport: cars,
	// }
	data := struct {
		Plant     []tree
		Transport []car
	}{
		trees,
		cars,
	}

	err := tmp.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalln(err)
	}

}
