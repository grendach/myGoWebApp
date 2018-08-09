package main

import (
	"log"
	"os"
	"text/template"
)

type hotel struct {
	Name, Address, City, Stars string
}

type region struct {
	Region string
	Hotels []hotel
}

type Regions []region

var tmp *template.Template

func init() {
	tmp = template.Must(template.ParseFiles("tmp.gohtml"))
}

func main() {
	reg := Regions{
		region{
			Region: "Obolon",
			Hotels: []hotel{
				hotel{
					Name:    "Obolon-Plaza",
					Address: "Gerojiv Stalingradu 18",
					City:    "Kyiv",
					Stars:   "***",
				},
				hotel{
					Name:    "U dupi",
					Address: "Dupska 77",
					City:    "Kyiv",
					Stars:   "*",
				},
			},
		},
		region{
			Region: "Podil",
			Hotels: []hotel{
				hotel{
					Name:    "VAL",
					Address: "V.Val 34",
					City:    "Kyiv",
					Stars:   "****",
				},
			},
		},
		region{
			Region: "Troya",
			Hotels: []hotel{
				hotel{
					Name:    "Capital",
					Address: "Kvitucha 84b",
					City:    "Kyiv",
					Stars:   "*****",
				},
			},
		},
	}

	err := tmp.Execute(os.Stdout, reg)
	if err != nil {
		log.Println(err)
	}
}
