package update

import (
	"log"
	"time"

	"github.com/spf13/cobra"

	"github.com/moohbr/WebMonitor/data"
	"github.com/moohbr/WebMonitor/infrastructure/database"
)

var (
	verbose bool

	UpdateCmd = &cobra.Command{
		Use:   "update",
		Short: "Update something",
		Long:  `With this command you can update something, like the servers or the users.`,
		Run: func(cmd *cobra.Command, args []string) {
			log.Println("[HELP] Update something use the subcommands")
			log.Println("[HELP] Use 'update server' to update a server")
			log.Println("[HELP] Use 'update user' to update a user")
		},
	}

	updateServerCmd = &cobra.Command{
		Use:   "server",
		Short: "Update a server",
		Long:  `With this command you can update a server.`,
		Run: func(cmd *cobra.Command, args []string) {
			log.Println("[SYSTEM] Updating a server")

			name, _ := cmd.Flags().GetString("name")
			ip, _ := cmd.Flags().GetString("ip")
			url, _ := cmd.Flags().GetString("url")
			monitor, _ := cmd.Flags().GetBool("monitor")

			server := data.Server{
				Name:       name,
				IP:         ip,
				URL:        url,
				Monitor:    monitor,
				LastUpdate: time.Now().UTC().Format("2006-01-02 15:04:05"),
			}

			db := database.OpenDatabase()
			db.UpdateServer(server)
		},
	}

	updateUserCmd = &cobra.Command{
		Use:   "user",
		Short: "Update a user",
		Long:  `With this command you can update a user.`,
		Run: func(cmd *cobra.Command, args []string) {
			log.Println("[SYSTEM] Updating a user")

			name, _ := cmd.Flags().GetString("name")
			password, _ := cmd.Flags().GetString("password")

			user := data.User{
				Name:     name,
				Password: password,
			}

			db := database.OpenDatabase()
			db.UpdateUser(user)
		},
	}
)

func init() {
	UpdateCmd.AddCommand(updateServerCmd)

	updateServerCmd.Flags().StringP("name", "n", "", "The name of the server")
	updateServerCmd.Flags().StringP("ip", "i", "", "The ip of the server")
	updateServerCmd.Flags().StringP("url", "u", "", "The url of the server")
	updateServerCmd.Flags().BoolP("monitor", "m", false, "If the server should be monitored")
	updateServerCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "Verbose output")

	UpdateCmd.AddCommand(updateUserCmd)
	updateUserCmd.Flags().StringP("name", "n", "", "The name of the user")
	updateUserCmd.Flags().StringP("password", "p", "", "The password of the user")
	updateUserCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "Verbose output")
}
