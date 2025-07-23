package migrate

import (
	"github.com/spf13/cobra"
	"honey_server/internal/common"
)

var logger = common.Logs().Cat("cmd/migrate")

func MigrateCommand() *cobra.Command {
	return migrateCmd
}

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "deal with migrate model",
	Long:  `this is a long description with migrate command`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		logger.Infoln("hello migrate")
	},
}
