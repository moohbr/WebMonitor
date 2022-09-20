package show

import (
	"fmt"
	"log"

	database "github.com/moohbr/WebMonitor/src/infrastructure/database"
	"github.com/spf13/cobra"
)

var (
	verbose bool
	ShowCmd = &cobra.Command{
		Use:   "show",
		Short: "Show something",
		Long:  `With this command you can show something, like the servers or the users.`,
		Run: func(cmd *cobra.Command, args []string) {
			db := database.NewDatabase()

			servers := db.GetServers()
			if len(servers) > 0 {

				for _, server := range servers {
					log.Println(server)
				}
			} else {
				log.Println("[SYSTEM] No servers found!")
			}
		},
	}

	showServerCmd = &cobra.Command{
		Use:   "server",
		Short: "Show a server",
		Long:  `With this command you can show a server.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Show a server")
		},
	}

	showServersCmd = &cobra.Command{
		Use:   "all",
		Short: "Show all servers",
		Long:  `With this command you can show all servers.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Show all servers")
		},
	}

	showUsersCmd = &cobra.Command{
		Use:   "users",
		Short: "Show all users",
		Long:  `With this command you can show all users.`,
		Run: func(cmd *cobra.Command, arg []string) {
			fmt.Println("Show all users")
		},
	}
)

func ShowInit() {
	ShowCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")
	ShowCmd.AddCommand(showServerCmd)
	ShowCmd.AddCommand(showUsersCmd)

	showServerCmd.AddCommand(showServersCmd)
}
