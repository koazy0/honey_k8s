package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"honey_server/internal/model"
)

type CreateLogReq struct {
	g.Meta `path:"/log" method:"post" tags:"日志管理" summary:"新增日志"`
	model.CreateLog
}

type CreateLogRes struct {
	model.CreateLogResponse
}

type ListLogReq struct {
	g.Meta `path:"/log" method:"get" tags:"日志管理" summary:"查看日志"`
	model.ListLog
}

type ListLogRes struct {
	model.ListLogResponse
}

type DeleteLogReq struct {
	g.Meta `path:"/log" method:"delete" tags:"主机管理" summary:"删除日志"`
	model.DeleteLog
}

type DeleteLogRes struct {
	model.DeleteLogResponse
}
