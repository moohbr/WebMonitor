package show

import (
	"fmt"
	"log"
	"os"
	"strconv"

	database "github.com/moohbr/WebMonitor/src/infrastructure/database"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

const (
	LayoutISO = "2006-01-02 15:04:05"
)

var (
	verbose bool
	ShowCmd = &cobra.Command{
		Use:   "show",
		Short: "Show something",
		Long:  `With this command you can show something, like the servers or the users.`,
		Run: func(cmd *cobra.Command, args []string) {
			log.Println("[SYSTEM] Add something use the subcommands")
			log.Println("[SYSTEM] Use 'show server' to show a server")
			log.Println("[SYSTEM] Use 'show user' to show a user")
		},
	}

	showServerCmd = &cobra.Command{
		Use:   "server",
		Short: "Show a server",
		Long:  `With this command you can show a server.`,
		Run: func(cmd *cobra.Command, args []string) {
			log.Println("[SYSTEM] Showing a server")
			db := database.NewDatabase()
			server := db.GetServer(args[0])
			if server.Name != "" {
				table := tablewriter.NewWriter(os.Stdout)
				table.SetHeader([]string{"Name", "IP", "URL", "Avarage Response Time", "Last Update", "Last Check", "Last Status", "Monitor"})
				laststatus := strconv.Itoa(server.LastStatus)
				table.Append([]string{server.Name, server.IP, server.URL, server.AvarageResponseTime.String(),
					server.LastUpdate.Format(LayoutISO), server.LastCheck.Format(LayoutISO), laststatus, strconv.FormatBool(server.Monitor)})
				defer table.Render()
			} else {
				log.Println("[SYSTEM] No server found!")
			}

		},
	}

	showServersCmd = &cobra.Command{
		Use:   "servers",
		Short: "Show all servers",
		Long:  `With this command you can show all servers.`,
		Run: func(cmd *cobra.Command, args []string) {
			db := database.NewDatabase()

			servers := db.GetServers()
			if len(servers) > 0 {
				table := tablewriter.NewWriter(os.Stdout)
				table.SetHeader([]string{"Name", "IP", "URL", "Avarage Response Time", "Last Update", "Last Check", "Last Status", "Monitor"})
				for _, server := range servers {
					laststatus := strconv.Itoa(server.LastStatus)
					table.Append([]string{server.Name, server.IP, server.URL, server.AvarageResponseTime.String(),
						server.LastUpdate.Format(LayoutISO), server.LastCheck.Format(LayoutISO), laststatus, strconv.FormatBool(server.Monitor)})

				}
				defer table.Render()
			} else {
				log.Println("[SYSTEM] No servers found!")
			}
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

func init() {
	ShowCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")
	ShowCmd.AddCommand(showServersCmd)
	ShowCmd.AddCommand(showServerCmd)

}
