package node

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"honey_server/api/node/v1"
)

func (c *ControllerV1) OptionsNode(ctx context.Context, req *v1.OptionsNodeReq) (res *v1.OptionsNodeRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
