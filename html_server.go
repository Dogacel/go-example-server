package main

import (
	"html/template"
	"net/http"
	"strconv"
)

//HTMLServer for serving a website
type HTMLServer struct {
	Port int
}

type todo struct {
	Title string
	Done  bool
}

type todoPageData struct {
	PageTitle string
	Todos     []todo
}

//Setup sets up the htmlserver
func (h *HTMLServer) Setup() {
	handleURLs()
}

//Start the html server
func (h *HTMLServer) Start() {
	http.ListenAndServe(":"+strconv.Itoa(h.Port), nil)
}

func handleURLs() {
	http.HandleFunc("/", indexHandler)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("pages/index.html"))
	data := todoPageData{
		PageTitle: "My TODO list",
		Todos: []todo{
			{Title: "Task 1", Done: false},
			{Title: "Task 2", Done: true},
			{Title: "Task 3", Done: true},
		},
	}
	tmpl.Execute(w, data)
}
