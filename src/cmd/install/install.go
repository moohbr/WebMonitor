package install

import (
	"log"

	"github.com/spf13/cobra"

	database "github.com/moohbr/WebMonitor/src/infrastructure/database"
)

var (
	verbose    bool
	InstallCmd = &cobra.Command{
		Use:   "install",
		Short: "Install database",
		Long:  `With this command you can install to storage data, like the servers or the users.`,
		Run: func(cmd *cobra.Command, args []string) {
			log.Println("[SYSTEM] Installing database")
			database.NewDatabase()
			log.Println("[SYSTEM] Database installed")
		},
	}
)

func init() {
	InstallCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")
}
