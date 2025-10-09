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
		return
	}

	// Déterminer quelle page afficher selon l'URL
	var templateFile string
	switch r.URL.Path {
	case "/easy":
		templateFile = "./page/easy.html"
	case "/normal":
		templateFile = "./page/normal.html"
	case "/hard":
		templateFile = "./page/hard.html"
	default:
		templateFile = "./page/normal.html"
	}

	template, err := template.ParseFiles(templateFile, "./template/header.html", "./template/footer.html")
	if err != nil {
		log.Fatal(err)
	}
	template.Execute(w, player)
}

func Difficulty(w http.ResponseWriter, r *http.Request, player *power4.Players) {
	if r.Method == "POST" {
		level := r.FormValue("level")

		// Ajouter un log pour déboguer
		log.Printf("Level received: '%s'", level)

		// Rediriger vers la page appropriée selon la difficulté choisie
		switch level {
		case "easy":
			log.Println("Redirecting to /easy")
			http.Redirect(w, r, "/easy", http.StatusSeeOther)
		case "normal":
			log.Println("Redirecting to /normal")
			http.Redirect(w, r, "/normal", http.StatusSeeOther)
		case "hard":
			log.Println("Redirecting to /hard")
			http.Redirect(w, r, "/hard", http.StatusSeeOther)
		default:
			log.Printf("Unknown level: '%s', redirecting to normal", level)
			http.Redirect(w, r, "/normal", http.StatusSeeOther)
		}
		return
	}

	// Afficher la page de sélection de difficulté
	template, err := template.ParseFiles("./page/difficulty.html", "./template/header.html", "./template/footer.html")
	if err != nil {
		log.Fatal(err)
	}
	template.Execute(w, nil)
}

func main() {
	var player power4.Players
	http.HandleFunc("/normal", func(w http.ResponseWriter, r *http.Request) {
		Player(w, r, &player)
	})
	http.HandleFunc("/easy", func(w http.ResponseWriter, r *http.Request) {
		Player(w, r, &player)
	})
	http.HandleFunc("/hard", func(w http.ResponseWriter, r *http.Request) {
		Player(w, r, &player)
	})
	http.HandleFunc("/", Home)
	http.HandleFunc("/difficulty", func(w http.ResponseWriter, r *http.Request) {
		Difficulty(w, r, &player)
	})
	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.ListenAndServe(":8080", nil)
}
