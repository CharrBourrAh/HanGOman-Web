package main

import (
	"fmt"
	"log"
	server "main/internal"
	"net/http"
)

func main() {
	log.Println("Starting our simple http server.")

	http.HandleFunc("/", server.Home)
	http.HandleFunc("/info", server.Info)

	log.Println("Started on port", server.PortNum)
	fmt.Println("To close connection CRLT+C")

	err := http.ListenAndServe(server.PortNum, nil)
	if err != nil {
		log.Fatal(err)
	}
}
