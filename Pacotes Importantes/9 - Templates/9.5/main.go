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
		templates := []string{
			"./header.html",
			"./content.html",
			"./footer.html",
		}

		users := []Usuario{
			{"Jack", "jack@golang.com"},
			{"Izzy", "izzy@golang.com"},
			{"Lino", "lino@golang.com"},
		}
		t := template.Must(template.New("content.html").ParseFiles(templates...))

		err := t.Execute(w, users)
		if err != nil {
			panic(err)
		}
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
