package main

import "net/http"

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/", HomeHandler)
	mux.Handle("/user", usuario{nome: "Jacksom"})

	http.ListenAndServe(":8080", mux)

}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}

type usuario struct {
	nome string
}

func (u usuario) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(u.nome))
}
