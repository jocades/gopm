package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

var opts struct {
	Host string `short:"h" long:"host" default:"127.0.0.1"`
	Port string `short:"p" long:"port"`
}

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8421")
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	fmt.Println("Connected to server")

	input := bufio.NewReader(os.Stdin)
	reader := bufio.NewReader(conn)

	for {
		// Read user input
		fmt.Print("Enter message: ")
		msg, err := input.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input: ", err)
			break
		}

		// Send message to server

		_, err = conn.Write([]byte(msg))
		if err != nil {
			fmt.Println("Error sending message to server: ", err)
			break
		}

		// Read server response
		res, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading server response: ", err)
			break
		}

		fmt.Printf("Server response: %s", res)
	}

}
