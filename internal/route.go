package internal

import (
	"net/http"
)

func RootHandler(server *http.ServeMux) {
	server.HandleFunc("/", Home)
	server.HandleFunc("/how-to-play", HowToPlay)
	server.HandleFunc("/game", Game)
	server.HandleFunc("/Win-Lose", WinLose)
}
