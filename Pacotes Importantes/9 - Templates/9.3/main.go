package main

import (
	"html/template"
	"os"
)

type Usuario struct {
	Nick  string
	Email string
}

func main() {

	users := []Usuario{
		{"Jack", "jack@golang.com"},
		{"Izzy", "izzy@golang.com"},
		{"Lino", "lino@golang.com"},
	}
	t := template.Must(template.New("template.html").ParseFiles("./template.html"))

	err := t.Execute(os.Stdout, users)
	if err != nil {
		panic(err)
	}
}
