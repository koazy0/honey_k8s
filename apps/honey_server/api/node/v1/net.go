package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"honey_server/internal/model"
)

//node相关的api

type ListNodeReq struct {
	g.Meta `path:"/node" method:"get" tags:"节点管理" summary:"查看节点"`
	model.ListNode
}

type ListNodeRes struct {
	model.ListNodeResponse
}

type UpdateNodeReq struct {
	g.Meta `path:"/node" method:"put" tags:"节点管理" summary:"更新节点"`
	model.UpdateNode
}

type UpdateNodeRes struct {
	model.UpdateNodeResponse
}

type DetailNodeReq struct {
	g.Meta `path:"/node/:id" method:"get" tags:"节点管理" summary:"查看特定节点详情"`
	model.DetailNode
}

type DetailNodeRes struct {
	model.DetailNodeResponse
}

type DeleteNodeReq struct {
	g.Meta `path:"/node/:id" method:"delete" tags:"节点管理" summary:"删除节点"`
	model.DeleteNode
}

type DeleteNodeRes struct {
	model.DeleteNodeResponse
}

type OptionsNodeReq struct {
	g.Meta `path:"/node/options" method:"get" tags:"节点管理" summary:"查看节点"`
	model.OptionsNode
}

type OptionsNodeRes struct {
	model.OptionsNodeResponse
}

//node网络相关的api

type ListNetworkNodeReq struct {
	g.Meta `path:"/node/network" method:"get" tags:"节点管理" summary:"查看节点网络"`
	model.ListNetworkNode
}

type ListNetworkNodeRes struct {
	model.ListNetworkNodeResponse
}

type UpdateNetworkNodeReq struct {
	g.Meta `path:"/node/network" method:"put" tags:"节点管理" summary:"更新节点网络"`
	model.UpdateNetworkNode
}

type UpdateNetworkNodeRes struct {
	model.UpdateNetworkNodeResponse
}

type FlushNetworkNodeReq struct {
	g.Meta `path:"/node/network/flush" method:"get" tags:"节点管理" summary:"刷新节点网络"`
	model.FlushNetworkNode
}

type FlushNetworkNodeRes struct {
	model.FlushNetworkNodeResponse
}

type EnableNetworkNodeReq struct {
	g.Meta `path:"/node/network/enable" method:"put" tags:"节点管理" summary:"启用节点网络"`
	model.EnableNetworkNode
}

type EnableNetworkNodeRes struct {
	model.EnableNetworkNodeResponse
}

type DeleteNetworkNodeReq struct {
	g.Meta `path:"/node/network/:id" method:"delete" tags:"节点管理" summary:"删除节点网络"`
	model.DeleteNetworkNode
}

type DeleteNetworkNodeRes struct {
	model.DeleteNetworkNodeResponse
}
