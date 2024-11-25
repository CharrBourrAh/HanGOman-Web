package internal

import (
	"fmt"
	"html/template"
	"log"
	"main/internal/hangman-classic/internal-hangman-classic/game"
	"main/internal/hangman-classic/pkg/structs"
	"net/http"
)

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
	t := template.Must(template.ParseGlob("./front-end/index.gohtml"))
	var data structs.HangManData
	err := t.Execute(w, nil)
	if err != nil {
		return
	}
	data.Nickname, data.WordFile = r.URL.Query().Get("nickname"), r.URL.Query().Get("word")
	game.Init(data.WordFile, &data, r)
	Game(w, r)
}

func HowToPlay(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseGlob("./front-end/how-to-play.gohtml"))
	err := t.Execute(w, nil)
	if err != nil {
		return
	}

}

func Game(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseGlob("./front-end/game.gohtml"))
	err := t.Execute(w, nil)
	if err != nil {
		return
	}
}
