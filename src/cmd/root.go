package cmd

import (
	"log"

	"github.com/spf13/cobra"

	show "github.com/moohbr/WebMonitor/src/cmd/show"
)

var (
	verbose bool

	RootCmd = &cobra.Command{
		Use:   "WebMonitor",
		Short: "WebMonitor is a tool to monitor websites",
		Long: `WebMonitor is a tool to monitor websites.
It will ping the websites and send a report by email.`,
		Run: func(cmd *cobra.Command, args []string) {
			log.Println("WebMonitor was developed by moohbr")
		},
	}
)

func init() {

	RootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")

	RootCmd.AddCommand(show.ShowCmd)
}
