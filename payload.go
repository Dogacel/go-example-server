package main

import (
	"bytes"
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

	// Create an *exec.Cmd
	cmd := exec.Command(".", "/webhook.sh")

	// Stdout buffer
	cmdOutput := &bytes.Buffer{}
	cmd.Stdout = cmdOutput

	err := cmd.Run() // will wait for command to return

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(cmdOutput) // => go version go1.3 darwin/amd64
}

func amain() {
	fmt.Println("Started !")
	http.HandleFunc("/payload", payloadHandler)
	http.ListenAndServe(":5050", nil)
}
