package add

import (
	"log"
	"time"

	"github.com/spf13/cobra"

	data "github.com/moohbr/WebMonitor/src/data"
	database "github.com/moohbr/WebMonitor/src/infrastructure/database"
)

var (
	verbose bool

	AddCmd = &cobra.Command{
		Use:   "add",
		Short: "Add something",
		Long:  `With this command you can add something, like the servers or the users.`,
		Run: func(cmd *cobra.Command, args []string) {
			log.Println("[SYSTEM] Add something use the subcommands")
			log.Println("[SYSTEM] Use 'add server' to add a server")
			log.Println("[SYSTEM] Use 'add user' to add a user")
		},
	}

	addServerCmd = &cobra.Command{
		Use:   "server",
		Short: "Add a server",
		Long:  `With this command you can add a server.`,
		Run: func(cmd *cobra.Command, args []string) {
			log.Println("[SYSTEM] Adding a server")
			name := args[0]
			url := args[1]
			server := data.Server{
				Name:                name,
				IP:                  url,
				AvarageResponseTime: 0,
				LastUpdate:          time.Now(),
				LastCheck:           time.Now(),
				LastStatus:          0,
				Monitor:             true,
			}

			db := database.NewDatabase()
			db.AddServer(server)
		},
	}

	addUserCmd = &cobra.Command{
		Use:   "user",
		Short: "Add a user",
		Long:  `With this command you can add a user.`,
		Run: func(cmd *cobra.Command, args []string) {
			log.Println("[SYSTEM] Add a user")
		},
	}
)

func init() {
	AddCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")
	AddCmd.AddCommand(addServerCmd)
	AddCmd.AddCommand(addUserCmd)
}
