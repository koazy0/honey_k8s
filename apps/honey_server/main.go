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
}
