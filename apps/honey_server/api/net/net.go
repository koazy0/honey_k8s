// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package net

import (
	"context"

	"honey_server/api/net/v1"
)

type INetV1 interface {
	CreateNet(ctx context.Context, req *v1.CreateNetReq) (res *v1.CreateNetRes, err error)
	ListNet(ctx context.Context, req *v1.ListNetReq) (res *v1.ListNetRes, err error)
	OptionsNet(ctx context.Context, req *v1.OptionsNetReq) (res *v1.OptionsNetRes, err error)
	DetailNet(ctx context.Context, req *v1.DetailNetReq) (res *v1.DetailNetRes, err error)
	UpdateNet(ctx context.Context, req *v1.UpdateNetReq) (res *v1.UpdateNetRes, err error)
	ScanNet(ctx context.Context, req *v1.ScanNetReq) (res *v1.ScanNetRes, err error)
	UseIPNet(ctx context.Context, req *v1.UseIPNetReq) (res *v1.UseIPNetRes, err error)
	DeleteNet(ctx context.Context, req *v1.DeleteNetReq) (res *v1.DeleteNetRes, err error)
}
