package main

import (
	"bufio"
	"fmt"
	"net"
	"os"

	"github.com/spf13/cobra"
)

var conn net.Conn
var reader *bufio.Reader
var writer *bufio.Writer

// instead of having the conn, reader and writer separate, create a Connection struct that holds them

type Connection struct {
	conn   net.Conn
	reader *bufio.Reader
	writer *bufio.Writer
}

// new connection function
func NewConnection(addr string) (*Connection, error) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return nil, err
	}

	return &Connection{
		conn:   conn,
		reader: bufio.NewReader(conn),
		writer: bufio.NewWriter(conn),
	}, nil
}

func (c *Connection) Send(msg string) error {
	_, err := c.writer.WriteString(msg)
	if err != nil {
		return err
	}
	c.writer.Flush()
	return nil
}

func (c *Connection) Close() {
	c.conn.Close()
}

var cli = &cobra.Command{
	Use:              "client",
	Short:            "A simple TCP client",
	Long:             "A simple TCP client to communicate with a server",
	PersistentPreRun: connect,
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		if conn != nil {
			conn.Close()
		}
	},
}

var ping = &cobra.Command{
	Use:   "ping",
	Short: "Ping the server",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Pinging the server")
		if conn == nil {
			fmt.Println("Connection not established")
			return
		}

		// Send a message to the server
		_, err := writer.WriteString("ping\n")
		if err != nil {
			fmt.Println("Error sending message to server: ", err)
			return
		}
		writer.Flush()

		// Read the server response
		res, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading server response: ", err)
			return
		}

		fmt.Printf("Server response: %s\n", res)
	},
}

func init() {
	cli.PersistentFlags().String("host", "127.0.0.1", "Host to connect to")
	cli.PersistentFlags().Int("port", 8421, "Port to connect to")

	cli.AddCommand(ping)
}

func connect(cmd *cobra.Command, args []string) {
	host, _ := cmd.Flags().GetString("host")
	port, _ := cmd.Flags().GetInt("port")
	addr := fmt.Sprintf("%s:%d", host, port)

	fmt.Printf("Connecting to %s...\n", addr)

	var err error
	conn, err = net.Dial("tcp", addr)
	if err != nil {
		fmt.Println("error connecting to server:", err)
		os.Exit(1)
	}
	reader = bufio.NewReader(conn)
	writer = bufio.NewWriter(conn)
}

func main() {
	if err := cli.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
