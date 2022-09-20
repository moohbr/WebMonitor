package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	ShowCmd = &cobra.Command{
		Use:   "show",
		Short: "Show something",
		Long:  `With this command you can show something, like the servers or the users.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Show all servers")
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
