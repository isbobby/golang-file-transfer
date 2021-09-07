package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	// listen for connection at port 8080
	ln, err := net.Listen("tcp", ":8080")

	fmt.Print("Listening for TCP connection at 8080\n")

	if err != nil {
		// handle error
		fmt.Print(err)
	}

	for {
		conn, err := ln.Accept()

		fmt.Printf("Accepting TCP connection: %v\n", conn)
		if err != nil {
			fmt.Print(err)
		}
		go handleConnection(conn)
		fmt.Printf("Connection closed: %v\n", conn)
	}
}

func handleConnection(c net.Conn) {
	// The defer ensures that the connection is closed no matter how the function exits.
	defer c.Close()
	buf := make([]byte, 4096)
	var file_size int

	for {
		n, err := c.Read(buf)
		if err != nil || n == 0 {
			fmt.Printf("Error in reading buffer:%v\n", err)
		} else {
			file_size = n
			break
		}
		n, err = c.Write(buf[0:n])
		if err != nil {
			fmt.Printf("Error in writing buffer:%v\n", err)
		} else {
			break
		}
	}

	// Decode the buffer
	text_string := string(buf[:file_size])
	fmt.Printf("Received file, size:%v \n", file_size)

	// Get path to save file
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	// Write file to the path obatained
	error := os.WriteFile(path+"/received", []byte(text_string), 0644)
	if err != nil {
		log.Println(error)
	}
	fmt.Printf("File saved under: " + path + "/received\n")
}
