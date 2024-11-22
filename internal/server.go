package internal

import (
	"fmt"
	"html/template"
	"log"
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
	t := template.Must(template.ParseGlob("./front-end/*.gohtml"))
	err := t.Execute(w, nil)
	if err != nil {
		return
	}

}

func Info(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Info page")
}
