package cmd

import (
	"fmt"

	"github.com/mrbeskin/intds/client"
	"github.com/spf13/cobra"
)

var port string
var host string
var clientGrpcPort string

// clientCmd represents the client command
var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "Start a client for the specified server.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("intds client starting . . . ")
		c := client.NewTcpClient(host, port, clientGrpcPort)
		fmt.Println("Dialing")
		c.DialTcp()
	},
}

func init() {
	rootCmd.AddCommand(clientCmd)
	clientCmd.Flags().StringVarP(&port, "port", "p", "", "Server Port to connect to. (Required)")
	clientCmd.Flags().StringVarP(&clientGrpcPort, "gport", "g", "50051", "Port for gRPC server to listen for connections.")
	clientCmd.Flags().StringVarP(&host, "host", "H", "", "Server Host to connect to. (Required)")
	clientCmd.MarkFlagRequired("port")
	clientCmd.MarkFlagRequired("host")
}
