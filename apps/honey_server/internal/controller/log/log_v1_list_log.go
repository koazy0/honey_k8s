package log

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"honey_server/api/log/v1"
)

func (c *ControllerV1) ListLog(ctx context.Context, req *v1.ListLogReq) (res *v1.ListLogRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
