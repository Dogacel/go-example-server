package main

import (
	"fmt"

	"./src/compactweb"
)

func main() {
	fmt.Println("Started app !")

	port := 9911

	exampleServer := compactweb.Server{
		Addr: "localhost",
		Port: port,
	}

	exampleClient := compactweb.Client{
		Addr: "localhost",
		Port: port,
	}

	go exampleServer.Listen()

	exampleClient.Bind()

	exampleClient.UpdateEverySecond()

}
