package user

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"honey_server/internal/model"
	"honey_server/internal/service"

	"github.com/gogf/gf/v2/errors/gerror"

	"honey_server/api/user/v1"
)

func (c *ControllerV1) Login(ctx context.Context, req *v1.LoginReq) (res *v1.LoginRes, err error) {
	if req.UserID == "" || req.Password == "" {
		return nil, gerror.New("用户名或密码不能为空")
	}
	in := model.UserSignIn{}
	gconv.Struct(req, &in)
	out, err := service.User().ValidateUser(ctx, in)
	gconv.Struct(out, &res)
	return
}
