package main

import (
	"fmt"
	"io"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {

	fmt.Print(r.Body)

	r.Header.Set(
		"Content-Type",
		"text/html",
	)

	io.WriteString(w,
		"Hello",
	)
}

func main() {
	fmt.Println("Started !")
	http.HandleFunc("/payload", handler)
	http.ListenAndServe(":5050", nil)
}
