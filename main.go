package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"strings"

	"github.com/themoderngeek/Redimension/service"
)

var redimension = *service.NewRedimension()

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error starting service", err.Error())
		return
	}
	defer listener.Close()
	fmt.Println("Service started on port 8080")
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection", err.Error())
			return
		}
		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
			} else {
				fmt.Println("Error reading message:", err)
			}
			return // Exit the handler when the connection is closed
		}

		message = strings.TrimSpace(message)

		parts := strings.SplitN(message, " ", 3)
		if len(parts) < 2 {
			conn.Write([]byte("Invalid command\n"))
			continue
		}
		command := parts[0]
		key := parts[1]
		switch command {
		case "SET":
			if len(parts) != 3 {
				conn.Write([]byte("Invalid SET command\n"))
				continue
			}
			value := parts[2]
			redimension.Set(key, value)
			conn.Write([]byte("OK\n"))
		case "GET":
			value, exists := redimension.Get(key)
			if !exists {
				conn.Write([]byte("Key not found\n"))
			} else {
				conn.Write([]byte(value + "\n"))
			}
		default:
			conn.Write([]byte("Invalid command\n"))
		}
	}
}
