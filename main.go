package main

import (
	"fmt"
	"html"
	"html/template"
	"net/http"
)

type Home struct {
	Path string
	Name string
	Color string
}

// main function runs the app
func main() {
	port := ":3000"

	fs := http.FileServer(http.Dir("./public"))
	http.Handle("/public/", http.StripPrefix("/public/", fs))

	http.HandleFunc("/", root)
	fmt.Printf("Starting server on port %s\n", port)
	panic(http.ListenAndServe(port, nil))

}

// root is the handler function, it tells me what to do when a user visits root
func root(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "text/html; charset=utf-8")
	color := r.URL.Query().Get("color")
	if color == "" {
		color = "black"
	}
	data := Home{
		Path: html.EscapeString(r.URL.Path),
		Name: "Grunde",
		Color: color,
	}
	t, _ := template.ParseFiles("templates/home.html")
	t.Execute(w, data)
	return
}
