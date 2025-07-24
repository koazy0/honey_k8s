package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"honey_server/internal/model"
)

type CreateHostReq struct {
	g.Meta `path:"/host" method:"post" tags:"主机管理" summary:"新增主机"`
	model.CreateHost
}

type CreateHostRes struct {
	model.CreateHostResponse
}

type ListHostReq struct {
	g.Meta `path:"/host" method:"get" tags:"主机管理" summary:"查看主机"`
	model.ListHost
}

type ListHostRes struct {
	model.ListHostResponse
}

type DeleteHostReq struct {
	g.Meta `path:"/host" method:"delete" tags:"主机管理" summary:"删除主机"`
	model.DeleteHost
}

type DeleteHostRes struct {
	model.DeleteHostResponse
}
