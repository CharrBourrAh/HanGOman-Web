package internal

import (
	"html/template"
	"net/http"
)

// var game GameState
var tmpl *template.Template

func Init() {
	tmpl = template.Must(template.ParseGlob("templates/*.html"))
}

func RootHandler(server *http.ServeMux) {
	server.HandleFunc("/", Home)
	server.HandleFunc("/how-to-play", HowToPlay)
	//server.HandleFunc("/guess", handleGuess)
	//server.HandleFunc("newgame", handleNewGame)
}
