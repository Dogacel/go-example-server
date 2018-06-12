package compactweb

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
)

//Server for handling connections.
type Server struct {
	Addr string
	Port int
}

//Listen the connected clients.
func (s *Server) Listen(done chan bool) {
	listener, err := net.Listen("tcp", s.Addr+":"+strconv.Itoa(s.Port))
	fmt.Println("Listening on " + s.Addr + ":" + strconv.Itoa(s.Port))
	if err != nil {
		panic("Error listening " + s.Addr + ":" + strconv.Itoa(s.Port))
	}
	defer listener.Close()

	done <- true

	for {
		conn, _ := listener.Accept()
		reader := bufio.NewReader(conn)
		for {
			message, err := reader.ReadString('\n')
			if err != nil {
				conn.Close()
				return
			}
			fmt.Print("Recieved : " + message)
		}
	}
}
