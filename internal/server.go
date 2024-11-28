package internal

import (
	"fmt"
	"html/template"
	"log"
	"main/front-end/leaderboard"
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
	log.Println("The http server is starting.")

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
			game.Init(context.data)
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
			println("input : " + context.data.Input)
			println("current word : " + context.data.ToFind)
			game.InsertChar(context.data)
			status = game.StatusGame(context.data)
			if status != "ingame" {
				switch status {
				case "win":
					context.data.Status = true
					leaderboard.AddPlayerToLeaderBoard(context.data.Nickname, context.data.Attempts, context.data.WordFile)
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
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			return
		}
		switch r.FormValue("endButtons") {
		case "replay":
			game.Init(context.data)
			http.Redirect(w, r, "/game", http.StatusSeeOther)
			return
		case "mainMenu":
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
	}
	t := template.Must(template.ParseGlob("./front-end/win-lose.gohtml"))
	err := t.Execute(w, context.data)
	if err != nil {
		return
	}
}

func LeaderBoardHandler(w http.ResponseWriter, r *http.Request) {
	boardData := leaderboard.GetLeaderBoard()

	tmpl, err := template.ParseFiles("./front-end/leaderboard.gohtml")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		http.Redirect(w, r, "/leaderboard", http.StatusSeeOther)
		return
	}
	err = tmpl.Execute(w, boardData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	return
}
