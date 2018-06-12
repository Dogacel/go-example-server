package main

import (
	"fmt"
	"io"
	"net/http"
	"os/exec"
)

func payloadHandler(w http.ResponseWriter, r *http.Request) {

	r.Header.Set(
		"Content-Type",
		"text/html",
	)

	io.WriteString(
		w,
		"Restarting server...",
	)

	exec.Command("cd ~/go-projects/go-example-server && git pull && go run main.go server")
}

func amain() {
	fmt.Println("Started !")
	http.HandleFunc("/payload", payloadHandler)
	http.ListenAndServe(":5050", nil)
}
