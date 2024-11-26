package internal

import (
	"fmt"
	"html/template"
	"log"
	"main/internal/hangman-classic/internal-hangman-classic/game"
	"main/internal/hangman-classic/pkg/structs"
	"net/http"
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
	} else if context.data.ToFind == "" {
		game.Init(context.data, r)
		println(context.data.ToFind)
	}
	/*
		for game.StatusGame(context.data) == "ingame" {
			if r.URL.Query().Get("word") != "" {
				context.data.Input = r.URL.Query().Get("word")
				game.InsertChar(context.data)
			}
		}
	*/
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			return
		}
		word := r.FormValue("word")
		println(word)
	}
	t := template.Must(template.ParseGlob("./front-end/game.gohtml"))
	err := t.Execute(w, *context.data)
	if err != nil {
		log.Println("Error executing game template:", err)
		return

	}

}
