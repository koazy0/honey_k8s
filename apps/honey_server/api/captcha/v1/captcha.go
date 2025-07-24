package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"honey_server/internal/model"
)

type GenerateCaptchaReq struct {
	g.Meta `path:"/captcha" method:"get" tags:"验证码" summary:"生产验证码图片"`
	model.GenerateCaptcha
}
type GenerateCaptchaRes struct {
	model.GenerateCaptchaReply
}
