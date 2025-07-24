// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package node

import (
	"context"

	"honey_server/api/node/v1"
)

type INodeV1 interface {
	ListNode(ctx context.Context, req *v1.ListNodeReq) (res *v1.ListNodeRes, err error)
	UpdateNode(ctx context.Context, req *v1.UpdateNodeReq) (res *v1.UpdateNodeRes, err error)
	DetailNode(ctx context.Context, req *v1.DetailNodeReq) (res *v1.DetailNodeRes, err error)
	DeleteNode(ctx context.Context, req *v1.DeleteNodeReq) (res *v1.DeleteNodeRes, err error)
	OptionsNode(ctx context.Context, req *v1.OptionsNodeReq) (res *v1.OptionsNodeRes, err error)
	ListNetworkNode(ctx context.Context, req *v1.ListNetworkNodeReq) (res *v1.ListNetworkNodeRes, err error)
	UpdateNetworkNode(ctx context.Context, req *v1.UpdateNetworkNodeReq) (res *v1.UpdateNetworkNodeRes, err error)
	FlushNetworkNode(ctx context.Context, req *v1.FlushNetworkNodeReq) (res *v1.FlushNetworkNodeRes, err error)
	EnableNetworkNode(ctx context.Context, req *v1.EnableNetworkNodeReq) (res *v1.EnableNetworkNodeRes, err error)
	DeleteNetworkNode(ctx context.Context, req *v1.DeleteNetworkNodeReq) (res *v1.DeleteNetworkNodeRes, err error)
}
