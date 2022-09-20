package show

import (
	"log"
	"os"
	"strconv"

	database "github.com/moohbr/WebMonitor/src/infrastructure/database"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

var (
	verbose bool
	ShowCmd = &cobra.Command{
		Use:   "show",
		Short: "Show something",
		Long:  `With this command you can show something, like the servers or the users.`,
		Run: func(cmd *cobra.Command, args []string) {
			log.Println("[HELP] Add something use the subcommands")
			log.Println("[HELP] Use 'show server' to show a server")
			log.Println("[HELP] Use 'show user' to show a user")
		},
	}

	showServerCmd = &cobra.Command{
		Use:   "server",
		Short: "Show a server",
		Long:  `With this command you can show a server.`,
		Run: func(cmd *cobra.Command, args []string) {
			if verbose {
				log.Println("[SYSTEM] Showing a server")
			}
			db := database.NewDatabase()
			server := db.GetServer(cmd.Flag("name").Value.String())
			if server.Name != "" {
				table := tablewriter.NewWriter(os.Stdout)
				table.SetHeader([]string{"Name", "IP", "URL", "Avarage Response Time", "Last Update", "Last Check", "Last Status", "Monitor"})
				laststatus := strconv.Itoa(server.LastStatus)
				table.Append([]string{server.Name, server.IP, server.URL, server.AvarageResponseTime.String(),
					server.LastUpdate, server.LastCheck, laststatus, strconv.FormatBool(server.Monitor)})
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
						server.LastUpdate, server.LastCheck, laststatus, strconv.FormatBool(server.Monitor)})
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
			db := database.NewDatabase()
			users := db.GetUsers()
			if len(users) > 0 {
				table := tablewriter.NewWriter(os.Stdout)
				table.SetHeader([]string{"Name", "Email", "Password", "Admin"})
				for _, user := range users {
					table.Append([]string{user.Name, user.Email, user.Password, strconv.FormatBool(user.Admin)})
				}
				defer table.Render()
			} else {
				log.Println("[SYSTEM] No users found!")
			}
		},
	}
)

func init() {
	ShowCmd.AddCommand(showServersCmd)
	showServerCmd.Flags().StringP("name", "n", "", "The name of the server")
	showServerCmd.MarkFlagRequired("name")
	showServerCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "Show more information")

	ShowCmd.AddCommand(showServerCmd)
	ShowCmd.AddCommand(showUsersCmd)
}
