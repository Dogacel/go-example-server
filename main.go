package main

import (
	"fmt"
	"os"

	"./compactweb"
)

func amain() {
	if len(os.Args) > 1 {
		typeFlag := os.Args[1]
		if typeFlag == "server" {
			fmt.Println("Server started !")
		} else if typeFlag == "client" {
			fmt.Println("Client started !")
		}
	}

	otherMain()
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
