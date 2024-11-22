package internal

import (
	"net/http"
)

func RootHandler(server *http.ServeMux) {
	server.HandleFunc("/", Home)
	server.HandleFunc("/info", Info)
	server.HandleFunc("/game", Game)
}
