package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"honey_server/internal/cmd/migrate"
	"honey_server/internal/cmd/server"
	"honey_server/internal/cmd/user"
	_ "honey_server/internal/cmd/user"
	"honey_server/internal/common"
	"os"
)

var logger = common.Logs().Cat("cmd")
var rootCmd = &cobra.Command{
	Use:   "honey",
	Short: "deal with user model",
	Long:  `this is a long description with user command`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		logger.Infoln("hello user")
	},
}

func init() {
	rootCmd.AddCommand(
		migrate.MigrateCommand(),
		server.ServerCommand(),
		user.UserCommand(),
	)
}
func RootCommand() *cobra.Command {
	return rootCmd
}

// Execute 对main函数暴露出来的方法
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

type FlagOptions struct {
	File    string
	Version bool
	DB      bool
	Menu    string
	Type    string
	Value   string
	Help    bool
}

var options FlagOptions

//func init() {
//	flag.StringVar(&options.File, "f", "settings.yaml", "配置文件路径")
//	flag.BoolVar(&options.Version, "vv", false, "打印当前版本")
//	flag.BoolVar(&options.Help, "h", false, "帮助信息")
//	flag.BoolVar(&options.DB, "db", false, "迁移表结构")
//	flag.StringVar(&options.Menu, "m", "", "菜单 user")
//	flag.StringVar(&options.Type, "t", "", "类型 create list")
//	flag.StringVar(&options.Value, "v", "", "值")
//	flag.Parse()
//	// 注册命令
//}
