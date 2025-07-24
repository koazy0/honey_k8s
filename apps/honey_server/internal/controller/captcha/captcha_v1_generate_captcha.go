package captcha

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"honey_server/api/captcha/v1"
)

func (c *ControllerV1) GenerateCaptcha(ctx context.Context, req *v1.GenerateCaptchaReq) (res *v1.GenerateCaptchaRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
