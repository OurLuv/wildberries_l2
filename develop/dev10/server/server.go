package server

import (
	"io"
	"log"
	"net"
)

func StartTCP() {
	addr := "localhost:8080"
	server, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalln(err)
	}
	defer server.Close()

	log.Println("Server is running on:", addr)

	for {
		conn, err := server.Accept()
		if err != nil {
			log.Println("Failed to accept conn.", err)
			continue
		}

		go func(conn net.Conn) {
			defer func() {
				conn.Close()
			}()
			io.Copy(conn, conn)
		}(conn)
	}
}
