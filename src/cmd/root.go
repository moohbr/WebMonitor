package cmd

import (
	"log"

	"github.com/spf13/cobra"

	add "github.com/moohbr/WebMonitor/src/cmd/add"
	install "github.com/moohbr/WebMonitor/src/cmd/install"
	ping "github.com/moohbr/WebMonitor/src/cmd/ping"
	remove "github.com/moohbr/WebMonitor/src/cmd/remove"
	show "github.com/moohbr/WebMonitor/src/cmd/show"
	update "github.com/moohbr/WebMonitor/src/cmd/update"
)

var (
	verbose bool

	RootCmd = &cobra.Command{
		Use:   "WebMonitor",
		Short: "WebMonitor is a tool to monitor websites",
		Long: `WebMonitor is a tool to monitor websites.
It will ping the websites and send a report by email.`,
		Run: func(cmd *cobra.Command, args []string) {
			log.Println("WebMonitor was developed by Matheus Araujo a.k.a. moohbr")
		},
	}
)

func init() {

	RootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")

	RootCmd.AddCommand(show.ShowCmd)
	RootCmd.AddCommand(add.AddCmd)
	RootCmd.AddCommand(update.UpdateCmd)
	RootCmd.AddCommand(remove.RemoveCmd)
	RootCmd.AddCommand(install.InstallCmd)
	RootCmd.AddCommand(ping.PingCmD)
}
