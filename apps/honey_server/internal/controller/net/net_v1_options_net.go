package net

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"honey_server/api/net/v1"
)

func (c *ControllerV1) OptionsNet(ctx context.Context, req *v1.OptionsNetReq) (res *v1.OptionsNetRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
