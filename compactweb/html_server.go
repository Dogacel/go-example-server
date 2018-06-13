package compactweb

import (
	"io"
	"net/http"
	"strconv"
)

//HTMLServer for serving a website
type HTMLServer struct {
	Port int
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
	r.Header.Set(
		"Content-Type",
		"text/html",
	)

	io.WriteString(
		w,
		"<h1> HTML Server is up and running ! </h1>",
	)
}
