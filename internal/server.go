package internal

import (
	"fmt"
	"net/http"
)

const PortNum string = ":3000"

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Homepage")
}

func Info(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Info page")
}
