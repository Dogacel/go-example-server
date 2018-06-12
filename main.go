package main

import (
	"fmt"

	"./compactweb"
)

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
