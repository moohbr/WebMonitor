package ping

import (
	"github.com/spf13/cobra"

	"github.com/moohbr/WebMonitor/src/agents/ping"
)

var (
	PingCmD = &cobra.Command{
		Use:   "ping",
		Short: "Ping a server",
		Long:  `With this command you can ping a server.`,
		Run: func(cmd *cobra.Command, args []string) {
			ping.PingAllServers()
		},
	}
)

func init() {

}
