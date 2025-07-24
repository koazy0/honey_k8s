package log

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"honey_server/api/log/v1"
)

func (c *ControllerV1) DeleteLog(ctx context.Context, req *v1.DeleteLogReq) (res *v1.DeleteLogRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
