package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"

	"./compactweb"
)

func main() {
	if len(os.Args) > 1 {
		typeFlag := os.Args[1]
		if typeFlag == "server" {
			fmt.Println("Server started !")
		} else if typeFlag == "client" {
			fmt.Println("Client started !")

			cmd := exec.Command("git", "commit -m 'Test' && git push origin master")

			cmdOutput := &bytes.Buffer{}
			cmd.Stdout = cmdOutput

			cmd.Start()
			cmd.Wait()

			fmt.Println(cmdOutput)
		}
	}
}

func otherMain() {
	fmt.Println("Started app !")

	port := 8081
	addr := "127.0.0.1"

	ok := make(chan bool)

	exampleServer := compactweb.Server{
		Addr: addr,
		Port: port,
	}

	exampleClient := compactweb.Client{
		Addr: addr,
		Port: port,
	}

	go exampleServer.Listen(ok)
	exampleClient.Bind(ok)

}
