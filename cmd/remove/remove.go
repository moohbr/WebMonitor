package remove

import (
	"log"

	"github.com/spf13/cobra"

	"github.com/moohbr/WebMonitor/infrastructure/database"
)

var (
	verbose bool

	RemoveCmd = &cobra.Command{
		Use:   "remove",
		Short: "Remove something",
		Long:  `With this command you can remove something, like the servers or the users.`,
		Run: func(cmd *cobra.Command, args []string) {
			log.Println("[HELP] Remove something use the subcommands")
			log.Println("[HELP] Use 'remove server' to remove a server")
			log.Println("[HELP] Use 'remove user' to remove a user")
		},
	}

	removeServerCmd = &cobra.Command{
		Use:   "server",
		Short: "Remove a server",
		Long:  `With this command you can remove a server.`,
		Run: func(cmd *cobra.Command, args []string) {
			log.Println("[SYSTEM] Removing a server")

			name, _ := cmd.Flags().GetString("name")
			if name == "" {
				log.Println("[ERROR] The name of the server is required")
				return
			}
			db := database.OpenDatabase()
			db.DeleteServer(name)
		},
	}

	removeUserCmd = &cobra.Command{
		Use:   "user",
		Short: "Remove a user",
		Long:  `With this command you can remove a user.`,
		Run: func(cmd *cobra.Command, args []string) {
			log.Println("[SYSTEM] Removing a user")

			name, _ := cmd.Flags().GetString("name")
			if name == "" {
				log.Println("[ERROR] The name of the user is required")
				return
			}
			db := database.OpenDatabase()
			db.DeleteUser(name)
		},
	}
)

func init() {
	RemoveCmd.AddCommand(removeServerCmd)
	removeServerCmd.PersistentFlags().StringP("name", "n", "", "The name of the server")
	err := removeServerCmd.MarkPersistentFlagRequired("name")

	if err != nil {
		log.Println("[ERROR] Error to mark the flag as required")
		log.Fatal(err)
	}

	removeServerCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "Verbose output")

	
	RemoveCmd.AddCommand(removeUserCmd)
	removeUserCmd.PersistentFlags().StringP("name", "n", "", "The name of the server")

	err = removeServerCmd.MarkPersistentFlagRequired("name")
	if err != nil {
		log.Println("[ERROR] Error to mark the flag as required")
		log.Fatal(err)
	}

	removeUserCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "Verbose output")
}
