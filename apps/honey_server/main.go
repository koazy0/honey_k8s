package main

import (
	_ "github.com/go-sql-driver/mysql"              // 必须匿名导入  database/sql 驱动
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2" // GoFrame ORM adapter
	"honey_server/internal/cmd"
	_ "honey_server/internal/logic"
	_ "honey_server/internal/packed" // 先进行匿名初始化
)

func main() {

	cmd.Execute()
	//var migrateFlag bool
	//ctx := gctx.GetInitCtx()
	//// 定义命令行参数
	//// 这里后面放在一个专门的包里面
	//flag.BoolVar(&migrateFlag, "migrate", false, "执行数据库迁移")
	//flag.BoolVar(&migrateFlag, "m", false, "执行数据库迁移（简写）")
	//flag.Parse()
	//
	//if migrateFlag {
	//	service.Migrations().Migrate(ctx)
	//}
	//cmd.Main.Run(ctx)

}
