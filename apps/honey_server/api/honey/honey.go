// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package honey

import (
	"context"

	"honey_server/api/honey/v1"
)

type IHoneyV1 interface {
	CreateHoneyIP(ctx context.Context, req *v1.CreateHoneyIPReq) (res *v1.CreateHoneyIPRes, err error)
	ListHoneyIP(ctx context.Context, req *v1.ListHoneyIPReq) (res *v1.ListHoneyIPRes, err error)
	DeleteHoneyIP(ctx context.Context, req *v1.DeleteHoneyIPReq) (res *v1.DeleteHoneyIPRes, err error)
	CreateHoneyPort(ctx context.Context, req *v1.CreateHoneyPortReq) (res *v1.CreateHoneyPortRes, err error)
	ListHoneyPort(ctx context.Context, req *v1.ListHoneyPortReq) (res *v1.ListHoneyPortRes, err error)
	DeleteHoneyPort(ctx context.Context, req *v1.DeleteHoneyPortReq) (res *v1.DeleteHoneyPortRes, err error)
}
