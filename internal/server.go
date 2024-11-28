package internal

import (
	"fmt"
	"html/template"
	"log"
	"main/internal/hangman-classic/internal-hangman-classic/game"
	"main/internal/hangman-classic/pkg/structs"
	"net/http"
	"strings"
)

type AppContext struct {
	data *structs.HangManData
}

var context AppContext

const PortNum string = ":3000"

func Init_Server() {
	router := http.NewServeMux()
	log.Println("Starting our simple http server.")

	RootHandler(router)
	log.Println("Started on port", PortNum)
	fmt.Println("To close connection CRLT+C")

	fs := http.FileServer(http.Dir("front-end"))
	router.Handle("/front-end/", http.StripPrefix("/front-end/", fs))

	err := http.ListenAndServe(PortNum, router)
	if err != nil {
		log.Fatal(err)
	}

}

func Home(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			return
		}
		data := &structs.HangManData{
			Nickname: r.FormValue("pseudo"),
			WordFile: r.FormValue("difficulty"),
		}
		if data.Nickname != "" && data.WordFile != "" {
			context.data = data
			if data.ToFind != "" {
				data.ToFind = ""
			}
			http.Redirect(w, r, "/game", http.StatusSeeOther)
			return
		}
	}
	t := template.Must(template.ParseGlob("./front-end/index.gohtml"))
	err := t.Execute(w, nil)
	if err != nil {
		return
	}

}

func HowToPlay(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseGlob("./front-end/how-to-play.gohtml"))
	err := t.Execute(w, nil)
	if err != nil {
		return
	}
}

func Game(w http.ResponseWriter, r *http.Request) {
	if context.data == nil {
		http.Error(w, "No hangman data", http.StatusBadRequest)
		return
	}
	if context.data.ToFind == "" {
		game.Init(context.data)
		println(context.data.ToFind)
	}
	status := game.StatusGame(context.data)
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			return
		}
		switch r.FormValue("pauseMenu") {
		case "replay":
			http.Redirect(w, r, "/game", http.StatusSeeOther)
			return
		case "mainMenu":
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
	}
	if status == "ingame" {
		if r.Method == "POST" {
			err := r.ParseForm()
			if err != nil {
				return
			}
			context.data.Input = strings.ToLower(r.FormValue("word"))
			println(context.data.Input)
			println(context.data.ToFind)
			game.InsertChar(context.data)
			status = game.StatusGame(context.data)
			if status != "ingame" {
				switch status {
				case "win":
					context.data.Status = true
				case "lose":
					context.data.Status = false
				}
				http.Redirect(w, r, "/Win-Lose", http.StatusSeeOther)
				return
			}
		}
	}
	t := template.Must(template.ParseGlob("./front-end/game.gohtml"))
	err := t.Execute(w, *context.data)
	if err != nil {
		log.Println("Error executing game template:", err)
		return
	}
}

func WinLose(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseGlob("./front-end/win-lose.gohtml"))
	err := t.Execute(w, context.data)
	if err != nil {
		return
	}
}

/* func LeaderBoard(w http.ResponseWriter, r *http.Request) {
	var LeaderboardList []structs.Score
	LeaderboardList = []structs.Score{
		{Nickname: "test", Attempts: 1, Difficulty: "facile"},
	}
	for _, player := range LeaderboardList {
		if player.Attempts > 0 {
			LeaderboardList = append(LeaderboardList, player)
		}
	}
	t, err := template.ParseGlob("./front-end/leaderboard.gohtml")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = t.Execute(w, LeaderboardList)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

*/
