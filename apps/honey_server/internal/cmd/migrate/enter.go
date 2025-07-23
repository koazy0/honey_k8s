package migrate

import (
	"github.com/spf13/cobra"
	"honey_server/internal/common"
	"honey_server/internal/service"
)

var logger = common.Logs().Cat("cmd/migrate")

func init() {
	//migrateCmd.Flags().BoolVarP()
}
func MigrateCommand() *cobra.Command {
	return migrateCmd
}

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "migrate tables",
	Long:  `migrate tables to initial databases`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		logger.Infoln("migrating tables...")
		service.Migrations().Migrate(cmd.Context())
		logger.Infoln("migrate tables success")
	},
}
