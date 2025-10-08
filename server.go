package main

import (
	"log"
	"net/http"
	"text/template"
)

func Home(w http.ResponseWriter, r *http.Request) {
	template, err := template.ParseFiles("./index.html", "./template/header.html", "./template/footer.html")
	if err != nil {
		log.Fatal(err)
	}
	template.Execute(w, nil)
}

func main() {
	http.HandleFunc("/", Home)
	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.ListenAndServe(":8080", nil)
}
