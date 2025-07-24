package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"honey_server/internal/model"
)

type CreateNetReq struct {
	g.Meta `path:"/net" method:"post" tags:"网络管理" summary:"新增网络"`
	model.CreateNet
}

type CreateNetRes struct {
	model.CreateNetResponse
}

type ListNetReq struct {
	g.Meta `path:"/net" method:"get" tags:"网络管理" summary:"查看网络"`
	model.ListNet
}

type ListNetRes struct {
	model.ListNetResponse
}

type OptionsNetReq struct {
	g.Meta `path:"/net/options" method:"get" tags:"网络管理" summary:"查看网络选项"`
	model.OptionsNet
}

type OptionsNetRes struct {
	model.OptionsNetResponse
}

type DetailNetReq struct {
	g.Meta `path:"/net/:id" method:"get" tags:"网络管理" summary:"查看特定网络详细信息"`
	model.DetailNet
}

type DetailNetRes struct {
	model.DetailNetResponse
}

type UpdateNetReq struct {
	g.Meta `path:"/net" method:"put" tags:"网络管理" summary:"更新网络"`
	model.UpdateNet
}

type UpdateNetRes struct {
	model.UpdateNetResponse
}

type ScanNetReq struct {
	g.Meta `path:"/net/scan" method:"post" tags:"网络管理" summary:"扫描网络"`
	model.ScanNet
}

type ScanNetRes struct {
	model.ScanNetResponse
}

type UseIPNetReq struct {
	g.Meta `path:"/net/ip_list" method:"get" tags:"网络管理" summary:"查询已使用的网络"`
	model.UseIPNet
}

type UseIPNetRes struct {
	model.UseIPNetResponse
}
type DeleteNetReq struct {
	g.Meta `path:"/net" method:"delete" tags:"网络管理" summary:"删除网络"`
	model.DeleteNet
}

type DeleteNetRes struct {
	model.DeleteNetResponse
}
