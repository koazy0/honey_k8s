package main

import (
	"flag"
	_ "github.com/go-sql-driver/mysql"              // 必须匿名导入！  database/sql 驱动
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2" // GoFrame ORM adapter
	_ "honey_node/internal/packed"
	"honey_node/internal/service"

	_ "honey_node/internal/logic"

	"github.com/gogf/gf/v2/os/gctx"

	"honey_node/internal/cmd"
)

func main() {
	var migrateFlag bool
	var singleFlag bool

	ctx := gctx.GetInitCtx()
	// 定义命令行参数
	// 这里后面放在一个专门的包里面
	flag.BoolVar(&migrateFlag, "migrate", false, "执行数据库迁移")
	flag.BoolVar(&migrateFlag, "m", false, "执行数据库迁移（简写）")
	flag.BoolVar(&singleFlag, "s", false, "执行数据库迁移（简写）")
	flag.BoolVar(&singleFlag, "single", false, "执行数据库迁移（简写）")
	flag.Parse()

	if migrateFlag {
		service.Migrations().Migrate(ctx)
	}
	cmd.Main.Run(ctx)

}
