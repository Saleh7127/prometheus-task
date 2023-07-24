package cmd

import (
	"fmt"
	"github.com/Saleh7127/prometheus-task/api"

	"github.com/spf13/cobra"
)

var (
	port = 7127
)
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
	and usage of using your command. For example:

	Cobra is a CLI library for Go that empowers applications.
	This application is a tool to generate the needed files
	to quickly create a Cobra application.`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("server called")
		api.Start(port)
	},
}

func init() {
	serverCmd.Flags().IntVarP(&port, "port", "p", 7127, "Default Port For HTTP Server")
	rootCmd.AddCommand(serverCmd)

}
