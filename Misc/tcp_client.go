package main

import (
	"fmt"
	"net"
)

func main() {
	// Listen
	// Accept
	// Handle -> Create thread
	dstream, err := net.Listen("tcp", ":8080")

	if err != nil {
		fmt.Print(err)
		return
	}
	defer dstream.Close()

	for {
		con, err := dstream.Accept()
		if err != nil {
			fmt.Print(err)
			return
		}
		go handle(con)
	}
}

func handle(con net.Conn) {
	for {
		buffer := make([]byte, 1024)
		lineToRead, err := con.Read(buffer)
		if err != nil {
			fmt.Print(err)
			return
		}
		fmt.Println(buffer[:lineToRead])
	}
	con.Close()
}