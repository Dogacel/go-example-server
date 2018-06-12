package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
)

func payloadHandler(w http.ResponseWriter, r *http.Request) {

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
	http.HandleFunc("/payload", payloadHandler)
	http.ListenAndServe(":5050", nil)
}
