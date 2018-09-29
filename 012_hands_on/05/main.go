package main

import (
	"log"
	"os"
	"text/template"
)

type item struct {
	Name, Descrip string
	Price         float64
}

type meal struct {
	Meal string
	Item []item
}

type menu []meal

var tmp *template.Template

func init() {
	tmp = template.Must(template.ParseFiles("tmp.gohtml"))
}

func main() {
	m := menu{
		meal{
			Meal: "Brecfast",
			Item: []item{
				item{
					Name:    "Яйця",
					Descrip: "Смачно",
					Price:   4.56,
				},
				item{
					Name:    "Borsch",
					Descrip: "Yuhuuu",
					Price:   5.11,
				},
				item{
					Name:    "Past",
					Descrip: "fuuu",
					Price:   2.56,
				},
			},
		},
		meal{
			Meal: "Lunch",
			Item: []item{
				item{
					Name:    "Dupa",
					Descrip: "Not very",
					Price:   1.16,
				},
				item{
					Name:    "varenyky",
					Descrip: "Norm",
					Price:   5.11,
				},
				item{
					Name:    "Pasta",
					Descrip: "not very good",
					Price:   2.56,
				},
			},
		},
		meal{
			Meal: "Dinner",
			Item: []item{
				item{
					Name:    "Keks",
					Descrip: "Good",
					Price:   11.16,
				},
				item{
					Name:    "Salo",
					Descrip: "Super",
					Price:   15.11,
				},
				item{
					Name:    "Potato",
					Descrip: "Nice, and not healthy",
					Price:   22.56,
				},
			},
		},
	}

	err := tmp.Execute(os.Stdout, m)
	if err != nil {
		log.Fatalln(err)
	}
}
