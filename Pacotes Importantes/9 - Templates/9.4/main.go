package main

import (
	"html/template"
	"log"
	"net/http"
)

type Usuario struct {
	Nick  string
	Email string
}

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		users := []Usuario{
			{"Jack", "jack@golang.com"},
			{"Izzy", "izzy@golang.com"},
			{"Lino", "lino@golang.com"},
		}
		t := template.Must(template.New("template.html").ParseFiles("./template.html"))

		err := t.Execute(w, users)
		if err != nil {
			panic(err)
		}
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
