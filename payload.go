package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {

	b, _ := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	rdr1 := ioutil.NopCloser(bytes.NewBuffer(b))
	log.Printf("BODY: %q", rdr1)

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
