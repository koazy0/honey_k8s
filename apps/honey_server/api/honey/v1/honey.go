package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"honey_server/internal/model"
)

type CreateHoneyIPReq struct {
	g.Meta `path:"/ip" method:"post" tags:"蜜罐管理" summary:"新增蜜罐IP"`
	model.CreateHoneyIP
}

type CreateHoneyIPRes struct {
	model.CreateHoneyIPResponse
}

type ListHoneyIPReq struct {
	g.Meta `path:"/ip" method:"get" tags:"蜜罐管理" summary:"查看蜜罐IP"`
	model.ListHoneyIP
}

type ListHoneyIPRes struct {
	model.ListHoneyIPResponse
}

type DeleteHoneyIPReq struct {
	g.Meta `path:"/ip" method:"del" tags:"蜜罐管理" summary:"删除蜜罐IP"`
	model.DeleteHoneyIP
}

type DeleteHoneyIPRes struct {
	model.DeleteHoneyIPResponse
}

type CreateHoneyPortReq struct {
	g.Meta `path:"/port" method:"post" tags:"蜜罐管理" summary:"新增蜜罐端口"`
	model.CreateHoneyPort
}

type CreateHoneyPortRes struct {
	model.CreateHoneyPortResponse
}

type ListHoneyPortReq struct {
	g.Meta `path:"/port" method:"get" tags:"蜜罐管理" summary:"查看蜜罐端口"`
	model.ListHoneyPort
}

type ListHoneyPortRes struct {
	model.ListHoneyPortResponse
}

type DeleteHoneyPortReq struct {
	g.Meta `path:"/port" method:"del" tags:"蜜罐管理" summary:"删除蜜罐端口"`
	model.DeleteHoneyPort
}

type DeleteHoneyPortRes struct {
	model.DeleteHoneyPortResponse
}
