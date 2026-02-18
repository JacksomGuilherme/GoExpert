package main

import (
	"html/template"
	"log"
	"net/http"
	"strings"
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
		t := template.New("content.html")
		t.Funcs(template.FuncMap{"ToUpper": strings.ToUpper})
		t = template.Must(t.ParseFiles(templates...))

		err := t.Execute(w, users)
		if err != nil {
			panic(err)
		}
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
