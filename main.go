package main

import (
	"encoding/json"
	"fmt"
	"html"
	"html/template"
	"log"
	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Home struct {
	Path  string
	Name  string
	Color string
	Deets []Deet
}

type Deet struct {
	gorm.Model
	Name string
	Email string
}

type App struct {
	DB *gorm.DB
}

// main function runs the app
func main() {
	port := ":3000"

	db, err := gorm.Open("sqlite3", "/tmp/gorm.db")
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()
	db.AutoMigrate(&Deet{})
	a := &App{DB: db}

	fs := http.FileServer(http.Dir("./public"))
	http.Handle("/public/", http.StripPrefix("/public/", fs))

	http.HandleFunc("/", a.root)
	http.HandleFunc("/makepost", a.makeyPosty)
	http.HandleFunc("/api/json", jsonResponse)
	fmt.Printf("Starting server on port %s\n", port)
	panic(http.ListenAndServe(port, nil))

}

// root is the handler function, it tells me what to do when a user visits root
func (a *App) root(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "text/html; charset=utf-8")
	color := r.URL.Query().Get("color")
	if color == "" {
		color = "black"
	}
	var Deets []Deet
	a.DB.Find(&Deets)
	data := Home{
		Path:  html.EscapeString(r.URL.Path),
		Name:  "Grunde",
		Color: color,
		Deets: Deets,
	}
	t, _ := template.ParseFiles("templates/home.html")
	t.Execute(w, data)
	return
}

func (a *App) makeyPosty(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		http.ServeFile(w, r, "templates/makeyposty.html")
	case "POST":
		r.ParseForm()
		fmt.Println(r.Form["name"][0])
		fmt.Println(r.Form["email"][0])
		a.DB.Create(&Deet{Name: r.Form["name"][0], Email: r.Form["email"][0]})
		http.Redirect(w, r, "/", http.StatusFound)
	default:
		http.NotFound(w, r)
	}
}

func jsonResponse(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	funObject := struct {
		Name            string
		Age             int            `json:"age"`
		FavoriteSubmits map[string]int // a map object whose keys are strings and values are integers
	}{
		Name: "Cooper Lee Joergens",
		Age:  0,
		FavoriteSubmits: map[string]int{
			"twitter": 10,
			"slack":   45,
			"Denver Posts comments section": 120,
		},
	}
	body, err := json.Marshal(funObject)
	if err != nil {
		log.Panicf("api: could not JSON encode response: %s", err)
	}
	w.Write(body)
	return
}
