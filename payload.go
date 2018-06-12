package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {

	b, _ := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	var f interface{}
	data := json.Unmarshal(b, &f)

	fmt.Print(data)

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
