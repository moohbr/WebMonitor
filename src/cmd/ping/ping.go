package add

import (
	"github.com/spf13/cobra"

	ping "github.com/moohbr/WebMonitor/src/use_cases/ping"
)

var (
	verbose bool

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
