package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	// first get absolute path to file
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	// read the file content and store in byte array
	dat, err := os.ReadFile(path + "/data")
	if err != nil {
		log.Println(err)
	}

	sent := false

	for true {
		// if file not sent
		if sent == false {
			host_addr := os.Getenv("RECEIVER_APP_SERVICE_HOST")
			host_port := os.Getenv("RECEIVER_APP_SERVICE_PORT")
			receiver_addr := host_addr + ":" + host_port

			// obtain the TCP address
			tcpAddr, err := net.ResolveTCPAddr("tcp", receiver_addr)
			if err != nil {
				log.Println(err)
				continue
			}

			fmt.Printf("Obtained receiver address: %v\n", tcpAddr)

			// open a TCP connection
			conn, err := net.DialTCP("tcp", nil, tcpAddr)
			if err != nil {
				log.Println(err)
				continue
			}

			fmt.Printf("Connected to receiving pod: %v\n", receiver_addr)
			conn.SetWriteBuffer(4096)

			// write over TCP connection
			n, err := conn.Write([]byte(dat))
			if err != nil {
				fmt.Println(err)
				continue
			}

			sent = true
			fmt.Printf("Message sent, byte array size: %v\n", n)
		}
	}
}
