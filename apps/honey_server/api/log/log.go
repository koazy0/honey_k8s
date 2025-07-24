// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package log

import (
	"context"

	"honey_server/api/log/v1"
)

type ILogV1 interface {
	CreateLog(ctx context.Context, req *v1.CreateLogReq) (res *v1.CreateLogRes, err error)
	ListLog(ctx context.Context, req *v1.ListLogReq) (res *v1.ListLogRes, err error)
	DeleteLog(ctx context.Context, req *v1.DeleteLogReq) (res *v1.DeleteLogRes, err error)
}
