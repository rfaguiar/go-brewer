package main

import (
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
	"path"
	"path/filepath"
)

func main() {
	r := mux.NewRouter().StrictSlash(true)

	staticDir := http.Dir(filepath.Join("src", "templates", "static"))
	r.PathPrefix("/static/").
		Handler(http.StripPrefix("/static/", http.FileServer(staticDir)))

	r.HandleFunc("/", ShowLogin).Methods(http.MethodGet)

	log.Print("Server startup and listen http://localhost:4000")
	log.Fatal(http.ListenAndServe(":4000", r))
}

func ShowLogin(w http.ResponseWriter, r *http.Request) {
	fp := path.Join("src", "templates", "login.html")
	tmpl, err := template.ParseFiles(fp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, "login"); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
