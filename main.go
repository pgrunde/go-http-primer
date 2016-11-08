package main

import (
	"fmt"
	"html"
	"html/template"
	"net/http"
)

type Home struct {
	Path  string
	Name  string
	Color string
}

// main function runs the app
func main() {
	port := ":3000"

	fs := http.FileServer(http.Dir("./public"))
	http.Handle("/public/", http.StripPrefix("/public/", fs))

	http.HandleFunc("/", root)
	http.HandleFunc("/makepost", makeyPosty)
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
		Path:  html.EscapeString(r.URL.Path),
		Name:  "Grunde",
		Color: color,
	}
	t, _ := template.ParseFiles("templates/home.html")
	t.Execute(w, data)
	return
}

func makeyPosty(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		http.ServeFile(w, r, "templates/makeyposty.html")
	case "POST":
		r.ParseForm()
		fmt.Println(r.Form["name"][0])
		fmt.Println(r.Form["email"][0])
		http.ServeFile(w, r, "templates/makeyposty.html")
	default:
		http.NotFound(w, r)
	}
}
