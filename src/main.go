package main

import (
	"html/template"
	"log"
	"net/http"
	"path"
	"path/filepath"
)

func main() {
	fs := http.FileServer(http.Dir(filepath.Join("src", "templates", "static")))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", ShowLogin)

	log.Print("Server startup and listen http://localhost:4000")
	log.Fatal(http.ListenAndServe(":4000", nil))
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
