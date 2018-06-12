package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
)

func handler(w http.ResponseWriter, r *http.Request) {

	requestDump, err := httputil.DumpRequest(r, true)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(requestDump))

	r.Header.Set(
		"Content-Type",
		"text/html",
	)

	io.WriteString(
		w,
		"Hello !",
	)
}

func main() {
	fmt.Println("Started !")
	http.HandleFunc("/payload", handler)
	http.ListenAndServe(":5050", nil)
}
