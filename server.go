package main

import (
	"log"
	"net/http"
	power4 "power4/src"
	"strconv"
	"text/template"
)

// Variable globale pour stocker le jeu en cours
var currentGame *power4.Game

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

	if currentGame != nil {
		data := struct {
			Player1       string
			Player2       string
			Player1_Score int
			Player2_Score int
			CurrentTurn   string
			Grid          [][]string
		}{
			Player1:       currentGame.Players.Player1,
			Player2:       currentGame.Players.Player2,
			Player1_Score: currentGame.Players.Player1_Score,
			Player2_Score: currentGame.Players.Player2_Score,
			CurrentTurn:   currentGame.Turn,
			Grid:          currentGame.Grid,
		}
		template.Execute(w, data)
	} else {
		template.Execute(w, player)
	}
}

func Difficulty(w http.ResponseWriter, r *http.Request, player *power4.Players) {
	if r.Method == "POST" {
		level := r.FormValue("level")
		log.Printf("Level received: '%s'", level)

		//récupérer la difficulté choisie
		player.Difficulty = level

		// Créer une nouvelle partie avec les joueurs et la difficulté
		currentGame = power4.NewGame(player)

		log.Printf("Game created: %+v", currentGame)

		switch level {
		case "easy":
			http.Redirect(w, r, "/easy", http.StatusSeeOther)
		case "normal":
			http.Redirect(w, r, "/normal", http.StatusSeeOther)
		case "hard":
			http.Redirect(w, r, "/hard", http.StatusSeeOther)
		default:
			http.Redirect(w, r, "/normal", http.StatusSeeOther)
		}
		return
	}

	template, err := template.ParseFiles("./page/difficulty.html", "./template/header.html", "./template/footer.html")
	if err != nil {
		log.Fatal(err)
	}
	template.Execute(w, nil)
}

// le placement des jetons
func PlaceTokenHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if currentGame == nil {
		log.Println("ERROR: No game in progress")
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	colStr := r.FormValue("col")
	col, err := strconv.Atoi(colStr)
	if err != nil {
		log.Printf("ERROR: Invalid column: %s", colStr)
		http.Error(w, "Invalid column", http.StatusBadRequest)
		return
	}

	log.Printf("Placing token in column %d", col)

	color := currentGame.GetCurrentPlayerColor()
	log.Printf("Current player: %s, color: %s", currentGame.Turn, color)

	success := currentGame.PlaceToken(col, color)

	if !success {
		log.Println("ERROR: Column full or invalid")
		redirectPath := "/" + currentGame.Players.Difficulty
		http.Redirect(w, r, redirectPath, http.StatusSeeOther)
		return
	}

	log.Println("Token placed successfully")

	// Vérifier la victoire AVANT de changer de joueur
	winner := currentGame.WinCond()
	if winner != "" {
		log.Printf("Winner detected: %s", winner)
		if winner == "red" {
			currentGame.Players.Player1_Score++
		} else {
			currentGame.Players.Player2_Score++
		}
		http.Redirect(w, r, "/win", http.StatusSeeOther)
		return
	}

	// Changer de joueur SEULEMENT si personne n'a gagné
	currentGame.SwitchTurn()
	log.Printf("Next turn: %s", currentGame.Turn)

	redirectPath := "/" + currentGame.Players.Difficulty
	log.Printf("Redirecting to %s", redirectPath)
	http.Redirect(w, r, redirectPath, http.StatusSeeOther)
}

// la page de victoire
func WinHandler(w http.ResponseWriter, r *http.Request) {
	if currentGame == nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	tmpl, err := template.ParseFiles("./page/win.html", "./template/header.html", "./template/footer.html")
	if err != nil {
		log.Fatal(err)
	}

	var winner string
	var winnerColor string
	if currentGame.Turn == currentGame.Players.Player1 {
		winner = currentGame.Players.Player1
		winnerColor = "red"
	} else {
		winner = currentGame.Players.Player2
		winnerColor = "yellow"
	}

	data := struct {
		Winner        string
		WinnerColor   string
		Player1       string
		Player2       string
		Player1_Score int
		Player2_Score int
		Difficulty    string
	}{
		Winner:        winner,
		WinnerColor:   winnerColor,
		Player1:       currentGame.Players.Player1,
		Player2:       currentGame.Players.Player2,
		Player1_Score: currentGame.Players.Player1_Score,
		Player2_Score: currentGame.Players.Player2_Score,
		Difficulty:    currentGame.Players.Difficulty,
	}

	tmpl.Execute(w, data)
}

func ReplayHandler(w http.ResponseWriter, r *http.Request) {
	if currentGame != nil {
		// Garder les joueurs et les scores
		currentGame = power4.NewGame(&currentGame.Players)
	}
	redirectPath := "/" + currentGame.Players.Difficulty
	http.Redirect(w, r, redirectPath, http.StatusSeeOther)
}

func main() {
	var player power4.Players

	http.HandleFunc("/", Home)
	http.HandleFunc("/difficulty", func(w http.ResponseWriter, r *http.Request) {
		Difficulty(w, r, &player)
	})
	http.HandleFunc("/normal", func(w http.ResponseWriter, r *http.Request) {
		Player(w, r, &player)
	})
	http.HandleFunc("/easy", func(w http.ResponseWriter, r *http.Request) {
		Player(w, r, &player)
	})
	http.HandleFunc("/hard", func(w http.ResponseWriter, r *http.Request) {
		Player(w, r, &player)
	})
	http.HandleFunc("/place-token", PlaceTokenHandler)
	http.HandleFunc("/win", WinHandler)
	http.HandleFunc("/replay", ReplayHandler)

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	log.Println("Server started on :8080")
	http.ListenAndServe(":8080", nil)
}
