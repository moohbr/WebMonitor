package add

import (
	"log"
	"strconv"
	"sync"
	"time"

	"github.com/spf13/cobra"

	data "github.com/moohbr/WebMonitor/src/data"
	database "github.com/moohbr/WebMonitor/src/infrastructure/database"
	templates "github.com/moohbr/WebMonitor/src/providers/mail/templates"
	mail "github.com/moohbr/WebMonitor/src/use_cases/sendMail"
)

var (
	wg      sync.WaitGroup
	verbose bool

	AddCmd = &cobra.Command{
		Use:   "add",
		Short: "Add something",
		Long:  `With this command you can add something, like the servers or the users.`,
		Run: func(cmd *cobra.Command, args []string) {
			log.Println("[HELP] Add something use the subcommands")
			log.Println("[HELP] Use 'add server' to add a server")
			log.Println("[HELP] Use 'add user' to add a user")
		},
	}

	addServerCmd = &cobra.Command{
		Use:   "server",
		Short: "Add a server",
		Long:  `With this command you can add a server.`,
		Run: func(cmd *cobra.Command, args []string) {
			defer wg.Wait()

			log.Println("[SYSTEM] Adding a server")
			name := cmd.Flag("name").Value.String()
			ip := cmd.Flag("ip").Value.String()
			url := cmd.Flag("url").Value.String()

			server := data.Server{
				Name:                name,
				IP:                  ip,
				URL:                 url,
				AvarageResponseTime: 0,
				LastUpdate:          time.Now().UTC().Format("2006-01-02 15:04:05"),
				LastCheck:           time.Now().UTC().Format("2006-01-02 15:04:05"),
				LastStatus:          0,
				Monitor:             true,
			}

			db := database.OpenDatabase()
			defer db.Close()
			db.AddServer(server)

			users := db.GetUsers()
			wg.Add(len(users))
			if len(users) > 0 {
				for _, user := range users {
					go mail.SendMail([]string{user.Email}, templates.NewServer(server), &wg)
				}
			} else {
				log.Println("[WARNING] No users to send the email")
			}

			log.Println("[SYSTEM] Server added")
		},
	}

	addUserCmd = &cobra.Command{
		Use:   "user",
		Short: "Add a user",
		Long:  `With this command you can add a user.`,
		Run: func(cmd *cobra.Command, args []string) {
			log.Println("[SYSTEM] Adding a user")
			name := cmd.Flag("name").Value.String()
			password := cmd.Flag("password").Value.String()
			email := cmd.Flag("email").Value.String()

			admin := cmd.Flag("admin").Value.String()

			adminBool, err := strconv.ParseBool(admin)
			if err != nil {
				log.Println("[ERROR] Error parsing the admin flag")
				log.Println("[ERROR] Error: ", err)
				return
			}

			user := data.User{
				Name:      name,
				Password:  password,
				Email:     email,
				Admin:     adminBool,
				LastLogin: time.Now(),
				LastNotif: time.Now(),
			}

			db := database.OpenDatabase()
			db.AddUser(user)
			defer db.Close()

			log.Println("[SYSTEM] User added")
		},
	}
)

func init() {
	AddCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")

	AddCmd.AddCommand(addServerCmd)
	addServerCmd.PersistentFlags().String("name", "", "The name of the server")
	addServerCmd.PersistentFlags().String("ip", "", "The ip of the server")
	addServerCmd.PersistentFlags().String("url", "", "The url of the server")
	addServerCmd.MarkPersistentFlagRequired("name")
	addServerCmd.MarkPersistentFlagRequired("ip")
	addServerCmd.MarkPersistentFlagRequired("url")
	addServerCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")

	AddCmd.AddCommand(addUserCmd)
	addUserCmd.PersistentFlags().String("name", "", "The name of the user")
	addUserCmd.PersistentFlags().String("password", "", "The password of the user")
	addUserCmd.PersistentFlags().String("email", "", "The email of the user")
	addUserCmd.PersistentFlags().Bool("admin", false, "The admin of the user")
	addUserCmd.MarkPersistentFlagRequired("name")
	addUserCmd.MarkPersistentFlagRequired("password")
	addUserCmd.MarkPersistentFlagRequired("email")
	addUserCmd.MarkPersistentFlagRequired("admin")
	addUserCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")
}
