package honey

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"honey_server/api/honey/v1"
)

func (c *ControllerV1) CreateHoneyPort(ctx context.Context, req *v1.CreateHoneyPortReq) (res *v1.CreateHoneyPortRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
