package remove

import (
	"log"

	"github.com/spf13/cobra"

	database "github.com/moohbr/WebMonitor/src/infrastructure/database"
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
			db := database.OpenDatabase()
			db.DeleteUser(name)
		},
	}
)

func init() {
	RemoveCmd.AddCommand(removeServerCmd)

	removeServerCmd.Flags().StringP("name", "n", "", "The name of the server")
	removeServerCmd.MarkPersistentFlagRequired("name")
	removeServerCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "Verbose output")

	RemoveCmd.AddCommand(removeUserCmd)
	removeUserCmd.Flags().StringP("name", "n", "", "The name of the user")
	removeServerCmd.MarkPersistentFlagRequired("name")
	removeUserCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "Verbose output")
}
