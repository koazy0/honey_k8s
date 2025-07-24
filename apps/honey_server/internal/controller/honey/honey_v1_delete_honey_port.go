package honey

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"honey_server/api/honey/v1"
)

func (c *ControllerV1) DeleteHoneyPort(ctx context.Context, req *v1.DeleteHoneyPortReq) (res *v1.DeleteHoneyPortRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
