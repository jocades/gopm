package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var cli = &cobra.Command{
	Use:   "client",
	Short: "A simple TCP client",
	Long:  "A simple TCP client to communicate with a server",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Client started")
		// print the args received
		fmt.Println("Args:", args)
		fmt.Println("Host:", cmd.Flag("host").Value)
		fmt.Println("Port:", cmd.Flag("port").Value)
	},
}

func Execute() {
	if err := cli.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cli.PersistentFlags().String("host", "127.0.0.1", "Host to connect to")
	cli.PersistentFlags().Int("port", 8421, "Port to connect to")
}
