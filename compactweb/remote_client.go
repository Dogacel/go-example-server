package compactweb

import (
	"fmt"
	"log"
	"net"
	"strconv"
	"time"
)

//Client for connecting to servers.
type Client struct {
	Addr string
	Port int
	conn net.Conn
}

//UpdateEverySecond Calls some function every second.
func (c *Client) UpdateEverySecond() {
	for count := 0; count <= 100; count++ {
		c.conn.Write([]byte(strconv.Itoa(count) + "\n"))
		fmt.Print("Sent " + strconv.Itoa(count) + "\n")
		time.Sleep(time.Second)
	}
}

//Bind to the server.
func (c *Client) Bind() {
	fmt.Println("Binding to " + c.Addr + ":" + strconv.Itoa(c.Port))
	conn, err := net.Dial("tcp", c.Addr+":"+strconv.Itoa(c.Port))
	if err != nil {
		log.Fatal(err)
		panic("Connection error !")
	}

	c.conn = conn
}
