package cmd

import (
	"fmt"

	"github.com/mrbeskin/intds/server"
	"github.com/spf13/cobra"
)

var serverPort string

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Starts an intds server listening on the specified port",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("intds server starting . . .")
		s := server.NewTcpServer(serverPort)
		s.ListenTcp()
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
	serverCmd.Flags().StringVarP(&serverPort, "port", "p", "", "Port on which to listen for connections. (Required)")
	serverCmd.MarkFlagRequired("port")
}
