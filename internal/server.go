package internal

import (
	"fmt"
	"html/template"
	"log"
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
	var data structs.HangManData
	context.data = &data
	//game.Init(&data, r)
	if r.URL.Query().Get("pseudo") != "" && r.URL.Query().Get("difficulty") != "" {
		data.Nickname = r.URL.Query().Get("pseudo")
		data.WordFile = r.URL.Query().Get("difficulty")
		http.Redirect(w, r, "/game", http.StatusSeeOther)
		return
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
	t := template.Must(template.ParseGlob("./front-end/game.gohtml"))
	var data structs.HangManData
	//game.Init(&data, r)
	err := t.Execute(w, data)
	if err != nil {
		return
	}
}
