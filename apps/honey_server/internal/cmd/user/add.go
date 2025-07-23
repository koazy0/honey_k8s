package user

import "github.com/spf13/cobra"

func init() {
	//解析命令行参数
	addCmd.Flags().BoolVarP(&add.Admin, "admin", "a", false, "add admin user")
	userCmd.AddCommand(addCmd)
}

type addOption struct {
	Admin bool //创建admin角色
}

var add addOption

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add users",
	Long:  `this is a long description with add command`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		logger.Infoln("hello add")
		logger.Infoln("args:", args)
		logger.Infof("addOption:%#v", add)

	},
}
