package net

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"honey_server/api/net/v1"
)

func (c *ControllerV1) ScanNet(ctx context.Context, req *v1.ScanNetReq) (res *v1.ScanNetRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
