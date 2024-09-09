package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"

	"github.com/spf13/cobra"
)

var cli = &cobra.Command{
	Use:   "server",
	Short: "A simple TCP server",
	Long:  "A simple TCP server to manage processes",
	Run:   serve,
}

func init() {
	cli.PersistentFlags().IntP("port", "p", 8421, "Port to listen on")
}

func main() {
	log.SetPrefix("PM: ")
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	if err := cli.Execute(); err != nil {
		log.Fatal(err)
	}
}

func serve(cmd *cobra.Command, args []string) {
	listener, err := net.Listen("tcp", ":8421")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Server listening on port", listener.Addr())

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		go handleConnection(conn)
	}
}

// Simple echo server
func handleConnection(conn net.Conn) {
	defer conn.Close()

	peer := conn.RemoteAddr()
	fmt.Println("Connection from", peer)

	reader := bufio.NewReader(conn)

	for {
		msg, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				log.Println("Connection closed by peer")
				break
			}
			log.Println("Error reading data from connection", err)
			break
		}

		fmt.Printf("Received: %s", msg)

		// Echo the message back to the client
		res := fmt.Sprintf("Echo: %s", msg)
		_, err = conn.Write([]byte(res))
		if err != nil {
			log.Println("Error writing data to connection", err)
			break
		}
	}

	fmt.Printf("Connection with %s closed.\n", peer)
}
