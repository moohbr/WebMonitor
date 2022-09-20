package add

import (
	"log"
	"strconv"
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
			ip := args[1]
			url := args[2]

			server := data.Server{
				Name:                name,
				IP:                  ip,
				URL:                 url,
				AvarageResponseTime: 0,
				LastUpdate:          time.Now(),
				LastCheck:           time.Now(),
				LastStatus:          0,
				Monitor:             true,
			}

			db := database.OpenDatabase()
			db.AddServer(server)
		},
	}

	addUserCmd = &cobra.Command{
		Use:   "user",
		Short: "Add a user",
		Long:  `With this command you can add a user.`,
		Run: func(cmd *cobra.Command, args []string) {
			log.Println("[SYSTEM] Adding a user")
			name := args[0]
			password := args[1]
			email := args[2]

			admin, err := strconv.ParseBool(args[3])
			if err != nil {
				log.Println("[ERROR] Error parsing admin")
				log.Println("[ERROR] Error: ", err)
			}

			user := data.User{
				Name:      name,
				Password:  password,
				Email:     email,
				Admin:     admin,
				LastLogin: time.Now(),
				LastNotif: time.Now(),
			}

			db := database.OpenDatabase()
			db.AddUser(user)
		},
	}
)

func init() {
	AddCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")
	AddCmd.AddCommand(addServerCmd)
	AddCmd.AddCommand(addUserCmd)
}
