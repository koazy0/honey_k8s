package host

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"honey_server/api/host/v1"
)

func (c *ControllerV1) ListHost(ctx context.Context, req *v1.ListHostReq) (res *v1.ListHostRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
