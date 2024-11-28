package internal

import (
	"fmt"
	"html/template"
	"log"
	"main/front-end/leaderboard"
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
	err := t.Execute(w, nil)
	if err != nil {
		return
	}
}

func LeaderBoardHandler(w http.ResponseWriter, r *http.Request) {
	leaderboard.AddPlayerToLeaderBoard("Alice", 10, "")
	leaderboard.AddPlayerToLeaderBoard("test", 5, "force a toi")
	leaderboard.AddPlayerToLeaderBoard("paul", 1, "facile")

	boardData := leaderboard.GetLeaderBoard()

	tmpl, err := template.ParseFiles("./front-end/leaderboard.gohtml")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, boardData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
