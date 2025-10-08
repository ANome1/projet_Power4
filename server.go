package main

import (
	"log"
	"net/http"
	power4 "power4/src"
	"text/template"
)

func Home(w http.ResponseWriter, r *http.Request) {
	template, err := template.ParseFiles("./index.html", "./template/header.html", "./template/footer.html")
	if err != nil {
		log.Fatal(err)
	}
	template.Execute(w, nil)
}

func Player(w http.ResponseWriter, r *http.Request, player *power4.Players) {
	if r.Method == "POST" {
		player.Player1 = r.FormValue("player1")
		player.Player2 = r.FormValue("player2")
		http.Redirect(w, r, "/difficulty", http.StatusSeeOther)
	}
	template, err := template.ParseFiles("./page/game.html", "./template/header.html", "./template/footer.html")
	if err != nil {
		log.Fatal(err)
	}
	template.Execute(w, player)
}

func Difficulty(w http.ResponseWriter, r *http.Request) {
	template, err := template.ParseFiles("./page/difficulty.html", "./template/header.html", "./template/footer.html")
	if err != nil {
		log.Fatal(err)
	}
	template.Execute(w, nil)
}
func main() {
	var player power4.Players
	http.HandleFunc("/game", func(w http.ResponseWriter, r *http.Request) {
		Player(w, r, &player)
	})
	http.HandleFunc("/", Home)
	http.HandleFunc("/difficulty", Difficulty)
	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.ListenAndServe(":8080", nil)
}
