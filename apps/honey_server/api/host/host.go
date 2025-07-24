// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package host

import (
	"context"

	"honey_server/api/host/v1"
)

type IHostV1 interface {
	CreateHost(ctx context.Context, req *v1.CreateHostReq) (res *v1.CreateHostRes, err error)
	ListHost(ctx context.Context, req *v1.ListHostReq) (res *v1.ListHostRes, err error)
	DeleteHost(ctx context.Context, req *v1.DeleteHostReq) (res *v1.DeleteHostRes, err error)
}
