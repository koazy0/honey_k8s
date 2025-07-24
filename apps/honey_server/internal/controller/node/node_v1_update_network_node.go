package node

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"honey_server/api/node/v1"
)

func (c *ControllerV1) UpdateNetworkNode(ctx context.Context, req *v1.UpdateNetworkNodeReq) (res *v1.UpdateNetworkNodeRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
