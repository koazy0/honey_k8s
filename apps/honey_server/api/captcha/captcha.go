// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package captcha

import (
	"context"

	"honey_server/api/captcha/v1"
)

type ICaptchaV1 interface {
	GenerateCaptcha(ctx context.Context, req *v1.GenerateCaptchaReq) (res *v1.GenerateCaptchaRes, err error)
}
