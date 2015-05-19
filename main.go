package main

import (
	"fmt"
	"net/http"
	"encoding/json"

	"./models"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!")
}

func Phones(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Server", "A Go Web Server")
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")

	ph := models.CreateModel()

	ph_Json, _ := json.Marshal(ph)
	fmt.Fprint(w, string(ph_Json))
}

func main() {
	http.HandleFunc("/", Index)
	http.HandleFunc("/phones", Phones)
	fmt.Println("Started on port 8008. Press Ctrl-C to stop.")
	http.ListenAndServe(":8008", nil)
}