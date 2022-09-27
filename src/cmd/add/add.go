package add

import (
	"log"
	"strconv"
	"time"

	"github.com/moohbr/WebMonitor/src/data"
	"github.com/moohbr/WebMonitor/src/infrastructure/database"
	"github.com/moohbr/WebMonitor/src/providers/mail/templates"
	"github.com/spf13/cobra"

	mail "github.com/moohbr/WebMonitor/src/agents/sendMail"
)

var (
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
			if len(users) > 0 {
				list_of_emails := []string{}
				for _, user := range users {
					list_of_emails = append(list_of_emails, user.Email)
				}
				mail.SendMail(list_of_emails, templates.ServerDown(server))
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

	err := addServerCmd.MarkPersistentFlagRequired("name")
	if err != nil {
		log.Println("[ERROR] Error marking the flag as required")
		log.Println("[ERROR] Error: ", err)
	}

	err = addServerCmd.MarkPersistentFlagRequired("ip")
	if err != nil {
		log.Println("[ERROR] Error marking the flag as required")
		log.Println("[ERROR] Error: ", err)
	}

	err = addServerCmd.MarkPersistentFlagRequired("url")
	if err != nil {
		log.Println("[ERROR] Error marking the flag as required")
		log.Println("[ERROR] Error: ", err)
	}

	addServerCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")

	AddCmd.AddCommand(addUserCmd)
	addUserCmd.PersistentFlags().String("name", "", "The name of the user")
	addUserCmd.PersistentFlags().String("password", "", "The password of the user")
	addUserCmd.PersistentFlags().String("email", "", "The email of the user")
	addUserCmd.PersistentFlags().Bool("admin", false, "The admin of the user")

	err = addUserCmd.MarkPersistentFlagRequired("name")
	if err != nil {
		log.Println("[ERROR] Error marking the flag as required")
		log.Println("[ERROR] Error: ", err)
	}

	err = addUserCmd.MarkPersistentFlagRequired("password")
	if err != nil {
		log.Println("[ERROR] Error marking the flag as required")
		log.Println("[ERROR] Error: ", err)
	}

	err = addUserCmd.MarkPersistentFlagRequired("email")
	if err != nil {
		log.Println("[ERROR] Error marking the flag as required")
		log.Println("[ERROR] Error: ", err)
	}

	err = addUserCmd.MarkPersistentFlagRequired("admin")
	if err != nil {
		log.Println("[ERROR] Error marking the flag as required")
		log.Println("[ERROR] Error: ", err)
	}
	addUserCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")
}
