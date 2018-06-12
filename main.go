package main

import (
	"fmt"

	"./compactweb"
)

func main() {
	fmt.Println("Started app !")

	port := 9911
	addr := "localhost"

	exampleServer := compactweb.Server{
		Addr: addr,
		Port: port,
	}

	exampleClient := compactweb.Client{
		Addr: addr,
		Port: port,
	}

	go exampleServer.Listen()

	exampleClient.Bind()

	exampleClient.UpdateEverySecond()

}
