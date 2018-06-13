package main

import (
	"fmt"
	"go-example-server/compactweb"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		typeFlag := os.Args[1]
		if typeFlag == "server" {
			fmt.Println("Server started !")
			serveHTML()
		} else if typeFlag == "client" {
			fmt.Println("Client started !")
		}
	}
}

func serveHTML() {
	server := compactweb.HTMLServer{Port: 8080}
	server.Setup()
	server.Start()
}

func testPorts() {
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
